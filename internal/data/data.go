package data

import (
	"echo-mangosteen/pkg/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(conf *config.Config) (*Data, func(), error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		conf.DB.Username,
		conf.DB.Passwd,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.DbName,
	)

	db, err := xorm.NewEngine(conf.DB.Driver, dsn)

	if err != nil {
		return nil, nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	db.ShowSQL(true)

	log.Print("connected to database")

	if err != nil {
		return nil, nil, err
	}

	log.Print("connected to redis")

	return &Data{DB: db}, func() {
		if err := db.Close(); err != nil {
			log.Errorf("failed to close database: %s", err)
		}
	}, nil
}
