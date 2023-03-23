package controller

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func Bye(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Bye, World!",
	})
}
