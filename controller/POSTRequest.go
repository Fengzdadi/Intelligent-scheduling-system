package controller

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system-back-end/common"
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
	// common.Insert(E)

	// 3. 返回请求的结果
	c.JSON(200, gin.H{
		"message": "Register success!",
	})
}

func DailySchedule(c *gin.Context) {
	// 请求天数据
	var dayOfWeek = c.PostForm("dayOfWeek")
	var s []string
	// 从数据库中读出
	switch dayOfWeek {
	case "1":
		s, err = common.ReadDailySchedule1(util.Pool)
	case "2":
		s, err = common.ReadDailySchedule2(util.Pool)
	case "3":
		s, err = common.ReadDailySchedule3(util.Pool)
	case "4":
		s, err = common.ReadDailySchedule4(util.Pool)
	case "5":
		s, err = common.ReadDailySchedule5(util.Pool)
	case "6":
		s, err = common.ReadDailySchedule6(util.Pool)
	case "7":
		s, err = common.ReadDailySchedule7(util.Pool)
	}
	// j := make(map[string]interface{})
	// err := json.Unmarshal(s, &j)
	// if err != nil {
	//	return
	//}
	if s != nil {
		c.JSON(200, s)
	} else {
		c.JSON(202, gin.H{
			"message": "Check ReadDailySchedule failed!",
		})
	}
}

func WeeklySchedule(c *gin.Context) {
	// 请求周数据

	var s []string
	// 从数据库中读出
	s, err = common.ReadWeeklySchedule(util.Pool)
	//j := make(map[string]interface{})
	//err := json.Unmarshal(s, &j)
	//if err != nil {
	//	return
	//}
	c.JSON(200, s)
}

func EnployeesAdd(c *gin.Context) {

	var employee model.Employee
	var emplike model.Emplike

	employee.Eid = c.PostForm("Eid")
	employee.Ename = c.PostForm("Ename")
	employee.Eemail = c.PostForm("Eemail")
	employee.Epos = c.PostForm("Epos")
	employee.Shid = c.PostForm("Shid")

	emplike.Eid = c.PostForm("Eid")
	emplike.Etype = c.PostForm("Etype")
	emplike.Elike = c.PostForm("Elike")

	var s string

	s, err = common.InsertEmployee(util.Pool, employee, emplike)
	if s == "success" {
		c.JSON(201, gin.H{
			"message": "employee add success!",
		})
	} else {
		c.JSON(202, gin.H{
			"message": "employee add failed!",
		})
	}
}

func EmployeeUpdate(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "employee update success!",
	})
}
