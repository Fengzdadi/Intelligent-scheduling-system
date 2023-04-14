package common

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-ini/ini"
	"log"
)

// 定义连接池结构体
type ConnPool struct {
	conns chan *sql.DB
}

// 初始化连接池
func NewConnPool(poolSize int) (*ConnPool, error) {
	conns := make(chan *sql.DB, poolSize)

	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Fatal(err)
	}

	conf := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		cfg.Section("sqlServer").Key("Data Source").Value(),
		cfg.Section("sqlServer").Key("user id").Value(),
		cfg.Section("sqlServer").Key("password").Value(),
		cfg.Section("sqlServer").Key("database").Value(),
	)

	for i := 0; i < poolSize; i++ {
		db, err := sql.Open("sqlserver", conf)
		//var db, err = sql.Open("sqlserver", strings.Join(conf, ";"))
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
