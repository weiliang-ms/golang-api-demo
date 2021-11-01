package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weiliang-ms/golang-api-demo/apis"
)

func main() {
	router := gin.Default()
	// 查询请求
	search := router.Group("/search")
	{
		search.GET("/timeout", apis.TimeoutByOneMinute)
	}
	router.Run(":8080")
}
