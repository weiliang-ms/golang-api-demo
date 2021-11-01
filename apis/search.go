package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func TimeoutByOneMinute(c *gin.Context) {
	log.Println("模拟1分钟业务逻辑...")
	time.Sleep(time.Minute)
	log.Println("业务处理完毕！")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
