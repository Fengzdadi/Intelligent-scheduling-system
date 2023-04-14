package common

import (
	"context"
	"github.com/go-ini/ini"
	_ "github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	"log"
	"net"
	"time"
)

type Config struct {
	RedisConf RedisConfig `ini:"Redis"`
}

type RedisConfig struct {
	Addr     string `ini:"Addr"`
	DB       int    `ini:"DB"`
	Username string `ini:"Username"`
	Password string `ini:"Password"`
	Network  int    `ini:"Network"`
}

var ctx = context.Background()

var DB *redis.Client

func RedisConnect() {
	// 连接redis
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Fatal()
		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Section("Redis").Key("Addr").Value(),
		DB:       cfg.Section("Redis").Key("DB").MustInt(0),
		Username: cfg.Section("Redis").Key("Username").Value(),
		Password: cfg.Section("Redis").Key("Password").Value(),
		Network:  cfg.Section("Redis").Key("Network").Value(),

		//连接池容量及闲置连接数量
		PoolSize:     15, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		// 可自定义连接函数
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			netDialer := &net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Minute,
			}
			return netDialer.Dial("tcp", "127.0.0.1:6379")
		},

		//钩子函数
		OnConnect: func(ctx context.Context, cn *redis.Conn) error { //仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
			log.Printf("conn=%v\n", cn)
			return nil
		},
	},
	)
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("连接成功")
	// 成功连接将其赋值给全局变量
	DB = rdb

}
