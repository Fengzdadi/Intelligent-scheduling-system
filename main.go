// This is the Go code for a simple login/logout handler

package main

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system-back-end/util"
)

func main() {

	util.DatabaseInit()

	r := gin.Default()

	r = CollectRoute(r)

	panic(r.Run())

}
