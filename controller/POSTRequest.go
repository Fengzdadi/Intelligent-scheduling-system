package controller

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system-back-end/commom"
	"github/Intelligent-scheduling-system-back-end/model"
	"github/Intelligent-scheduling-system-back-end/util"
)

var err error

func HandleRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func RegisterRequest(c *gin.Context) {
	// 将请求写入数据库

	// 1. 获取请求的参数
	var E model.Employee
	E.Eid = c.PostForm("Eid")
	E.Ename = c.PostForm("Ename")
	E.Eemail = c.PostForm("Eemail")
	E.Epos = c.PostForm("Epos")
	E.Shid = c.PostForm("Shid")

	// 2. 将参数写入数据库
	commom.Insert(E)

	// 3. 返回请求的结果
	c.JSON(200, gin.H{
		"message": "Register success!",
	})
}

func DailySchedule(c *gin.Context) {
	// 请求天数据

	dayOfWeek := c.PostForm("dayOfWeek")
	var s []byte
	// 从数据库中读出
	s, err = commom.ReadDailySchedule(util.Pool, dayOfWeek)
	c.JSON(200, s)
}

func WeeklySchedule(c *gin.Context) {
	// 请求周数据

	week := c.PostForm("week")
	var s []byte
	// 从数据库中读出
	s, err = commom.ReadWeeklySchedule(util.Pool)
	c.JSON(200, s)
}
