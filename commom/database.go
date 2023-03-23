package commom

//连接sql sever
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github/Intelligent-scheduling-system-back-end/model"
	"log"
	"strings"
	"time"
)

// 初始化数据库
var db *sql.DB

func TestSqlServer() {
	var conf []string
	var _ *sql.DB
	var err error
	conf = append(conf, "Provider=SQLSEVERDB")
	conf = append(conf, "Data Source=47.120.3.158") // sqlserver IP 和 服务器名称
	conf = append(conf, "Initial Catalog=test")     // 数据库名
	conf = append(conf, "user id=SA")               // 登陆用户名
	conf = append(conf, "password=Mzh1234@")        // 登陆密码

	fmt.Println(strings.Join(conf, ";"))
	_, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}

	// 连接数据库
	//db, err = sql.Open("adodb", strings.Join(conf, ";"))
	//if err != nil {
	//	fmt.Println("sql open:", err)
	//	return
	//}
	//// 插入数据
	//insert, err := db.Exec("INSERT INTO student (Name,Age,Score) VALUES (?, ?, ?)", "tom", 18, 88.5)
	//if err != nil {
	//	fmt.Println("Insert data err=", err)
	//}
	//fmt.Println("成功插入数据!")
}

// 连接池类型
type Pool struct {
	MaxOpenConns int
	MaxIdleConns int
}

func NewPool(maxOpenConns int) *Pool {
	return &Pool{MaxOpenConns: maxOpenConns}
}

// 连接池代码
func ConnectPool() {
	var conf []string
	var err error
	conf = append(conf, "Provider=SQLSEVERDB")
	conf = append(conf, "Data Source=47.120.3.158") // sqlserver IP 和 服务器名称
	conf = append(conf, "Initial Catalog=Test")     // 数据库名
	conf = append(conf, "user id=SA")               // 登陆用户名
	conf = append(conf, "password=Mzh1234@")        // 登陆密码
	db, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sql open:", err)
	return
}

func Insert(Employee model.Employee) {
	// 插入数据
	Attach()
	fmt.Println(strings.Join(conf, ";"))
	_, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}

	_, err = db.Exec("INSERT INTO Employee (Eid,Ename,Eemail,Epos,Shid) VALUES (?, ?, ?, ?, ?)", Employee.Eid, Employee.Ename, Employee.Eemail, Employee.Epos, Employee.Shid)
	if err != nil {
		fmt.Println("Insert data err=", err)
	}
	fmt.Println("成功插入数据!")
}

func InsertData(dbPool *commom.ConnPool) {
	// 获取连接
	db, err := dbPool.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Release(db)

	// 执行数据库操作
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句
	_, err = db.ExecContext(ctx, "INSERT INTO PaiBan (Name,Time,Classroom,Week,Section) VALUES (?, ?, ?, ?, ?)", "tom", "2021-05-11", "A102", "星期一", "第一节")
	if err != nil {
		log.Fatal(err)
	}
}

func ReadDailySchedule(dbPool *ConnPool, day string) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf("SELECT Planid, Eid, Epos, %s, FROM PaiBan", day)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	jsonData, err := json.Marshal(rows)

	return jsonData, nil
}

func ReadWeeklySchedule(dbPool *ConnPool) ([]byte, error) {
	db, err := dbPool.Get()
	if err != nil {
		return nil, err
	}
	defer dbPool.Release(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// sql查询语句，注意指定需要查询的列

	tsql := fmt.Sprintf("SELECT Planid, Eid, Epos, W1, W2, W3, W4, W5, W6, W7 FROM PaiBan")
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		fmt.Println("Error reading rows: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil { // 检查行集是否为空，防止空指针异常
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)

	jsonData, err := json.Marshal(rows)

	return jsonData, nil
}
