package tool

import (
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	engine, err := xorm.NewEngine(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return nil, err
	}
	return &Orm{engine}, nil
}
