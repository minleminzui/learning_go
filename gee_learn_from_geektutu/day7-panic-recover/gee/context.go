package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 使用Fprintf来向Wiriter写数据
	Writer http.ResponseWriter
	// 请求信息包装在Request中
	Req *http.Request
	// 请求路径，从Request中获取
	Path string
	// 请求方法，一般是GET或者POST，从Request中获取
	Method string
	// 解析后的动态路由Path， 比如`:`与`*`
	Params map[string]string
	// 状态码
	StatusCode int
	// 中间件列表
	handlers []HandlerFunc
	// Next方法中会使用到,初始值为-1
	index int
	// engine pointer
	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// 调用中间件的执行
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	// 最后的c.index++是要保证可以离开这个函数
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// 在服务器错误的时候返回500，并使得之后的中间件不再执行
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// 动态路由参数`:`与`*`
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 从POST请求的表单中获取数据
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 从GET方法的URL中获得数据
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 用来设置状态码 Status
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 用来设置`首部字段`
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// Response是String
func (c *Context) String(code int, format string, values ...interface{}) {
	// 此处设置实体首部字段
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Response是二进制数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// Response是JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	// 通过http.ResponseWriter来获取一个json encoder
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		// http包回应request，错误信息err.Error()和错误码code
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Reponse是HTML
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		// 这里为啥不用http.Error(c.Writer, err.Error(), 500)
		c.Fail(500, err.Error())
	}
}
