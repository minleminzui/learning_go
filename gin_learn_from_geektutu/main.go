package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1.生成了一个实例，也就是一个WSGI
	r := gin.Default()

	// 2.声明了一个路由，告诉Gin什么样的URL能触发传入
	// 的函数，这个函数用来返回我们想要显示在用户浏览器中
	// 的信息
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, Geektutu")
	// })

	// :name解析路径参数
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello, %s", name)
	// })

	// 匹配users?name=xxx&role=xxx, role可选
	// r.GET("/users", func(c *gin.Context) {
	// 	name := c.Query("name")
	// 	role := c.DefaultQuery("role", "teacher")
	// 	c.String(http.StatusOK, "%s is a %s", name, role)
	// })

	// POST 解析参数
	// r.POST("/form", func(c *gin.Context) {
	// 	username := c.PostForm("username")
	// 	password := c.DefaultPostForm("password", "000000") // 设置默认值

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// })

	// GET与POST混合
	// r.POST("/posts", func(c *gin.Context) {
	// 	id := c.Query("id")
	// 	page := c.DefaultQuery("page", "0")
	// 	username := c.PostForm("username")
	// 	password := c.DefaultPostForm("password", "000000")

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"id":       id,
	// 		"page":     page,
	// 		"usernmae": username,
	// 		"password": password,
	// 	})
	// })

	// POST中的map
	// r.POST("/post", func(c *gin.Context) {
	// 	ids := c.QueryMap("ids")
	// 	names := c.PostFormMap("names")

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"ids":   ids,
	// 		"names": names,
	// 	})
	// })

	// 重定向
	// r.GET("/redirect", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/index")
	// })

	// r.GET("/goindex", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/"
	// 	r.HandleContext(c)
	// })

	// 分组路由
	// defaultHandler := func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"path": c.FullPath(),
	// 	})
	// }

	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/posts", defaultHandler)
	// 	v1.GET("/series", defaultHandler)
	// }

	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/posts", defaultHandler)
	// 	v2.GET("/series", defaultHandler)
	// }

	// 上传文件
	// 单个文件
	// r.POST("/upload1", func(c *gin.Context) {
	// 	file, _ := c.FormFile("file")
	// 	// c.SaveUpload(file, dst)
	// 	c.String(http.StatusOK, "%s uploaded!", file.Filename)
	// })

	// // 多个文件
	// r.POST("/upload2", func(c *gin.Context) {
	// 	// Mulipart form
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["upload[]"]

	// 	for _, file := range files {
	// 		log.Println(file.Filename)
	// 		// c.SaveUploadedFile(file, dst)
	// 	}

	// 	c.String(http.StatusOK, "%d files uploaded!", len(files))
	// })

	// HTML模版(Template)
	// type student struct {
	// 	Name string
	// 	Age  int8
	// }

	// r.LoadHTMLGlob("templates/*")

	// stu1 := &student{Name: "Geektutu", Age: 20}
	// stu2 := &student{Name: "Jack", Age: 22}
	// r.GET("/arr", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "arr.tmpl", gin.H{
	// 		"title":  "Gin",
	// 		"stuArr": [2]*student{stu1, stu2},
	// 	})
	// })

	// Middleware(中间件)
	// 作用于全局
	// r.Use(gin.Logger())
	// r.Use(gin.Recvoery())

	// // 作用于单个路由
	// r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// // 作用于某个组
	// authorized := r.Group("/")
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("login", loginEndpoint)
	// 	authorized.POST("submit", submitEndpoint)
	// }
	// 3.要应用运行在本地服务器上，默认监听端口是
	// 8080,可以传入参数设置端口，比如r.Run(":9999")
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}
