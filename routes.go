package main

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system-back-end/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// Request
	// GET
	r.GET("/Hello", controller.Hello)
	r.GET("/ByeBye", controller.Bye)
	r.GET("/StartSchedule", controller.StartSchedule)

	// POST
	r.POST("/HandleRequest", controller.HandleRequest)
	r.POST("/RegisterRequest", controller.RegisterRequest)

	//r.POST("/DefaultValue", controller.DefaultValue)
	//r.POST("/EmployeeManagement", controller.Manageement)
	r.POST("/DailySchedule", controller.DailySchedule)
	r.POST("/WeeklySchedule", controller.WeeklySchedule)

	r.POST("/EmployeesAdd", controller.EnployeesAdd)
	return r
}
