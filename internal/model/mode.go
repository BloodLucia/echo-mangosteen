package model

import (
	"database/sql"
	"time"

	"github.com/spf13/cast"
)

type Model struct {
	ID        uint64       `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time    `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time    `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt sql.NullTime `xorm:"TIMESTAMP deleted_at"`
}

func (m *Model) GetStringID() string {
	return cast.ToString(m.ID)
}
