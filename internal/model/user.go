package model

import (
	"time"

	"github.com/spf13/cast"
)

type User struct {
	ID        int64     `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	Email     string    `xorm:"not null VARCHAR(100) UNIQUE email"`
}

type UserLoginRequest struct {
	Email string `validate:"required|email" message:"无效的邮箱地址" label:"用户邮箱" json:"email"`
	Code  string `validate:"required|min_len:6|max_len:6" message:"验证码长度应为 6" json:"code"`
}

type UserLoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type UserSendValidationCodeRequest struct {
	Email string `validate:"required|email" message:"无效的邮箱地址" label:"用户邮箱" json:"email"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) GetStringID() string {
	return cast.ToString(u.ID)
}
