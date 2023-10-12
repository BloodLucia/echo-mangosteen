package model

import (
	"database/sql"
	"time"

	"github.com/spf13/cast"
)

type Kind string

type Item struct {
	ID         uint64       `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt  time.Time    `xorm:"created TIMESTAMP created_at"`
	UpdatedAt  time.Time    `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt  sql.NullTime `xorm:"TIMESTAMP deleted_at"`
	UserId     uint64       `xorm:"not null BIGINT(20) user_id"`
	TagId      uint64       `xorm:"not null BIGINT(20) tag_id"`
	Kind       Kind         `xorm:"not null VARCHAR(10) kind"`
	Amount     int          `xorm:"not null DECIMAL(2) amount"`
	HappenedAt time.Time    `xorm:"not null DATETIME happened_at"`
}

type CreateItemRequest struct {
	TagId      string    `json:"tag_id"`
	Kind       string    `json:"kind"`
	Amount     int       `json:"amount"`
	HappenedAt time.Time `json:"happened_at"`
}

func (i *Item) TableName() string {
	return "items"
}

func (i *Item) GetStringID() string {
	return cast.ToString(i.ID)
}
