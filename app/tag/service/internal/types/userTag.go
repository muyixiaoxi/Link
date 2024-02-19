package types

import "gorm.io/gorm"

type UserTagFollow struct {
	gorm.Model
	TagId  uint64 `json:"tag_id"`
	UserId uint64 `json:"user_id"`
}

func (u *UserTagFollow) TableName() string {
	return "tb_user_tag"
}
