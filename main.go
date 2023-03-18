// This is the Go code for a simple login/logout handler

package main

import (
	"github.com/gin-gonic/gin"
)

func handleRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

// 在本地8080端口上运行一个html文件

func main() {
	router := gin.Default()

	router.GET("/", handleRequest)
	//对一个POST请求返回

	router.Run(":8080")

}
