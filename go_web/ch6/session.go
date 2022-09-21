// 1.cookie
// 	(1) 简而言之，cookie可以理解为在本地计算机保存一些用户操作的历史信息
// 并在用户再次访问浏览器时通过HTTP协议将本地cookie内容发送给服务器
// 从而完成验证
//	(2) cookie是由浏览器维持，存储在客户端的一小段文本信息，伴随着用户请求和页面
// 在Web服务器和浏览器之间传递，用户每次访问站点时，Web应用程序可以读取cookie包含的
// 的信息
//		会话cookie:生命周期从创建浏览器到关闭浏览器为止，会话cookie一般保存在内存中
//		持久cookie:保存在硬盘中
// 2. session
// 	就是在服务器上保存用户操作的历史信息，服务器使用session id来标识session，
// session id 由服务器负责产生，保证随机性与唯一性
// 指一类用来在客户端与服务端之间保持状态的解决方案

package main