// This is the Go code for a simple login/logout handler

package main

import (
	"github.com/gin-gonic/gin"
	"github/Intelligent-scheduling-system-back-end/util"
)

func main() {
	var err error
	_, err = util.DatabaseInit()
	if err != nil {
		return
	}

	r := gin.Default()

	r = CollectRoute(r)

	panic(r.Run())

}
