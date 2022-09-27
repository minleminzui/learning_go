>该项目学习自[极客兔兔教程](https://geektutu.com/post/gee.html),意在实现一个Go语言实现的Web框架Gee

## 序言
### net/http
- 标准库net/http只是提供了基础的Web功能
    - 监听端口
    - 映射静态路由
    - 解析HTTP报文
- 但是并没有提供类似    
    - 动态路由：诸如hello/:name, hello/*这种规则
    - 鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现
    - 模板：没有统一简化HTML机制
## day1
- 实现了
    - 自定义路由映射表
    - 提供了用户注册静态路由的方法
    - 包装了启动服务的函数
## day2
- 实现了
    - 将路由(router)独立出来
    - 设计上下文(Context), 封装Request和Reponse，提供对JSON，HTML等返回类型的支持。设计Context的方法
        - 封装了*http.Request 和 http.ResponseWriter方法，简化了相关接口的调用
        - 解析动态路由/hello/:name
        - 存储中间件信息
## day3
