package model

type Tag struct {
	Model
	UserId uint64 `xorm:"not null BIGINT(20) user_id"`
	Name   string `xorm:"not null varchar(50) name"`
	Type   string `xorm:"not null VARCHAR(10) type"`
	Sign   string `xorm:"not null CHAR(1) sign"`
}

type TagAddRequest struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Sign string `json:"sign" validate:"required"`
}

func (t *Tag) TableName() string {
	return "tags"
}
