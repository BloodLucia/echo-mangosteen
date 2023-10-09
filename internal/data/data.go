package data

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
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

	return &Data{DB: db}, func() {
		if err := db.Close(); err != nil {
			log.Errorf("failed to close db: %s", err)
		}
	}, nil
}
