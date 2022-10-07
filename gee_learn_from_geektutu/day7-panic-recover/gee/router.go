package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 用于将路径解析为数组
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			// *通配符只能作为最后一个
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	// 这里可以直接通过key来执行处理函数
	r.handlers[key] = handler
}

// 把node节点存在trie树中，只是为了判断相应路径是否存在处理函数
// 如果存在，可以直接从router的handlers中获取
func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)
	Params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// 这里通过寻找node，判断该路径是否是有效的
	n := root.search(searchParts, 0)
	if n != nil {
		// 注意这里需要解析n.pattern，而不是直接使用searchParts
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				Params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' {
				Params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, Params
	}

	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		// 注意这里要取的是n.pattern，其中存储的是带通配符的路径
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		// 如果对应路径不存在node，那么添加一个返回404的方法
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
