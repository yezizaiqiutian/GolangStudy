package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	test := r.Group("test")
	{
		test.GET("/test1", test1)
	}
	r.Run(":8888")
}

func test1(context *gin.Context) {
	name := context.DefaultQuery("name", "111")
	text := "以后咱们开个公司:\n"

	switch name {
	case "陈志家":
		text += "家家负责我们的后台程序,"
		break
	case "贾学敏":
		text += "贾哥负责把我们的产品卖出去,"
		break
	case "高辉":
		text += "高辉负责我们的app程序,"
		break
	}
	context.JSON(http.StatusOK, text)
}
