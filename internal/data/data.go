package data

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type Data struct {
	DB    *xorm.Engine
	Cache *redis.Client
}

func NewData() (*Data, func(), error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		"root",
		"root",
		"localhost",
		3306,
		"echo_mangosteen",
	)

	db, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		return nil, nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	log.Print("connected to database")

	rdb, err := NewRedis()
	if err != nil {
		return nil, nil, err
	}

	log.Print("connected to redis")

	return &Data{DB: db, Cache: rdb}, func() {
		if err := db.Close(); err != nil {
			log.Errorf("failed to close database: %s", err)
		}
		if err := rdb.Close(); err != nil {
			log.Errorf("failed to close redis: %s", err)
		}
	}, nil
}
