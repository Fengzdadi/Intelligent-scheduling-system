package controller

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system-back-end/common"
)

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

func StartSchedule(c *gin.Context) {
	_, err := common.AlgorithmExec()
	if err != nil {
		c.JSON(200, gin.H{
			"message": "StartSchedule, Success!",
		})
		return
	}

}
