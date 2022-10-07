package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// 定义处理request的函数
type HandlerFunc func(*Context)

// 定义Engine，它实现了ServerHTTP接口, 所以可以处理http request
type (
	RouterGroup struct {
		// 通过request的路径前缀来区分group
		prefix string
		// 这个group的中间件数组
		middlewares []HandlerFunc
		engine      *Engine // 所有group共用一个Engine实例
	}
	Engine struct {
		// 这里是go的嵌套类型，相当于继承，
		// Engine充当最顶层的group
		// 单纯的group只是负责分组路由
		// 而Engine有更多功能，比如Run，ServeHTTP
		*RouterGroup
		router        *router
		groups        []*RouterGroup     // 存储所有路由分组
		htmlTemplates *template.Template // html 渲染器
		funcMap       template.FuncMap   // html渲染器
	}
)

// Engine的构造器
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 添加了默认中间件Logger()和Recovery的Engine构造器
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

// 用来添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// 用来添加分组Group
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		engine: engine,
		prefix: group.prefix + prefix,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 用来添加新的路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// 添加GET方法的路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// 添加POST方法的路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// 创建static handler，用来将某个映射文件夹
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 获得要映射的文件夹的绝对路径
	absolutPath := path.Join(group.prefix, relativePath)
	// StripPrefix函数将absolutePath这个前缀去除
	// 相当于在路由器中，StripPrefix("/assets", http.FileServer(http.Dir("/usr/geektutu/blog/static"))
	// 相当于把前缀改成"/usr/geektutu/blog/static"
	fileServer := http.StripPrefix(absolutPath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// 暴露给用户的接口
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	// 还需要添加路由函数，因为每次request时，会通过方法+路径把handler加到这次
	// 请求的中间件列表中，之后去执行中间件以及路由函数
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}

// 定制化的渲染函数
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

// 用来将所有的html模板载入内存
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

// 包装了启动函数
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	// 这里加入所有相关group的中间件
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, req)
	c.engine = engine
	c.handlers = middlewares
	engine.router.handle(c)
}
