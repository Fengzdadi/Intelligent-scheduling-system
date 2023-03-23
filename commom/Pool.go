package commom

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

// 定义连接池结构体
type ConnPool struct {
	conns chan *sql.DB
}

// 初始化连接池
func NewConnPool(poolSize int) (*ConnPool, error) {
	conns := make(chan *sql.DB, poolSize)
	var conf []string
	conf = append(conf, "Provider=SQLSEVERDB")
	conf = append(conf, "Data Source=47.120.3.158") // sqlserver IP 和 服务器名称
	conf = append(conf, "Initial Catalog=Test")     // 数据库名
	conf = append(conf, "user id=SA")               // 登陆用户名
	conf = append(conf, "password=Mzh1234@")        // 登陆密码

	for i := 0; i < poolSize; i++ {
		db, err := sql.Open("adodb", strings.Join(conf, ";"))
		if err != nil {
			return nil, err
		}
		conns <- db
	}
	return &ConnPool{conns: conns}, nil
}

// 从连接池中获取一个连接
func (p *ConnPool) Get() (*sql.DB, error) {
	select {
	case db := <-p.conns:
		if err := db.Ping(); err != nil {
			return nil, err
		}
		return db, nil
	default:
		return nil, fmt.Errorf("connection pool is full")
	}
}

// 将连接放回连接池中
func (p *ConnPool) Release(db *sql.DB) {
	p.conns <- db
}

func main() {
	// 初始化连接池
	pool, err := NewConnPool(10)
	if err != nil {
		log.Fatal(err)
	}

	// 获取连接
	db, err := pool.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Release(db)

	// 执行数据库操作
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM myTable")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理查询结果
	for rows.Next() {
		// ...
	}
}
