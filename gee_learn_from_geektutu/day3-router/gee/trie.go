package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由， 例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点， 例如[doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part含有:或*为true
	// isWild参数，比如当匹配`/p/go/doc/`这个路由时
	// 第一层节点，`p`精确匹配到了`p`,`go`模糊匹配到
	// `:lang`,那么会把`lang`这个参数赋值为`go`,继续下一层匹配
}

// 任何实现String接口的类型都可以被Print调用
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// 第一匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		// 可以是`*`或则`:`通配
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
