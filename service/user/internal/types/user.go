package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:20;not null" json:"username"`
	Password string `gorm:"size:60;not null" json:"password"`
	Age      uint   `json:"age"`
	Gender   uint   `gorm:"comment:0: 未知 1：男 2：女" json:"gender"`
	Address  string `json:"address"`
	Phone    string `gorm:"size:11;" json:"phone"`
	Avatar   string `json:"avatar"`
}
