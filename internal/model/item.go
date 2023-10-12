package model

import (
	"database/sql"
	"time"

	"github.com/spf13/cast"
)

type Item struct {
	ID        uint64       `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time    `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time    `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt sql.NullTime `xorm:"TIMESTAMP deleted_at"`
	TagId     uint64
	UserId    uint64
	// TODO
}

func (i *Item) TableName() string {
	return "items"
}

func (i *Item) GetStringID() string {
	return cast.ToString(i.ID)
}
