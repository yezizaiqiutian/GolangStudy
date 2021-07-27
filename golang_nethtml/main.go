package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

//Gin框架结合packr打包html资源为可执行文件
//https://zhuanlan.zhihu.com/p/345127425
func main() {

	//本地运行可以,不能打包
	//r := gin.Default()
	//r.LoadHTMLGlob("golang_nethtml/tem/index.html")
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
	//})
	////r.Run()
	//r.Run(":8888")

	box := packr.NewBox("./tem")
	//映射静态资源文件
	r := gin.Default()
	r.StaticFS("/web", box)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8888")
}
