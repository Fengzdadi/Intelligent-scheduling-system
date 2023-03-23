package util

import (
	"github/Intelligent-scheduling-system-back-end/commom"
	"log"
)

var Pool *commom.ConnPool

// DatabaseInit 初始化数据库
func DatabaseInit() {
	// 初始化连接池
	var err error
	Pool, err = commom.NewConnPool(10)
	if err != nil {
		log.Fatal(err)
	}
}
