package main

//连接sql sever
import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-adodb"
)

const (
	userName = "SA"
	password = "Mzh1234@"
	ip       = "47.120.3.158"
	port     = "1433"
	dbName   = "test"
)

func TestSqlServer(BillMonth string, BusinessCodeName string, RealTotalCost float64) {
	var conf []string
	var db *sql.DB
	var err error
	conf = append(conf, "Provider=SQLSEVERDB")
	conf = append(conf, "Data Source=47.120.3.158") // sqlserver IP 和 服务器名称
	conf = append(conf, "Initial Catalog=test")     // 数据库名
	conf = append(conf, "user id=SA")               // 登陆用户名
	conf = append(conf, "password=Mzh1234@")        // 登陆密码
	//fmt.Println(strings.Join(conf, ";"))

	// 连接数据库
	db, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	// 插入数据
	insert, err := db.Exec("INSERT INTO student (Name,Age,Score) VALUES (?, ?, ?)", "tom", 18, 88.5)
	if err != nil {
		fmt.Println("Insert data err=", err)
	}
	fmt.Println("成功插入数据!")
}
