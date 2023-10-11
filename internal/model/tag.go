package model

import (
	"database/sql"
	"time"

	"github.com/spf13/cast"
)

type Tag struct {
	ID        uint64       `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time    `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time    `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt sql.NullTime `xorm:"TIMESTAMP deleted_at"`
	UserId    uint64       `xorm:"not null BIGINT(20) user_id"`
	Name      string       `xorm:"not null varchar(50) name"`
	Type      string       `xorm:"not null VARCHAR(10) type"`
	Sign      []byte       `xorm:"not null CHAR(1) sign"`
}

type TagAddRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Sign []byte `json:"sign"`
}

func (t *Tag) TableName() string {
	return "tags"
}

func (t *Tag) GetStringID() string {
	return cast.ToString(t.ID)
}
