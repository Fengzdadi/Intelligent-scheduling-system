package util

import (
	"github/Intelligent-scheduling-system-back-end/common"
	"log"
)

var Pool *common.ConnPool

// DatabaseInit 初始化数据库
func DatabaseInit() (string, error) {
	// 初始化连接池
	var err error
	Pool, err = common.NewConnPool(10)
	if err != nil {
		log.Fatal(err)
		return "failed", err
	}
	return "success", err
}
