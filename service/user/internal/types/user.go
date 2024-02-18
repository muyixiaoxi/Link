package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"size:20;not null"`
	Password   string `gorm:"size:60;not null" `
	Age        uint   `json:"age"`
	Gender     uint   `gorm:"comment:0: 未知 1：男 2：女"`
	Address    string `json:"address"`
	Phone      string `gorm:"size:11;" json:"phone"`
	Avatar     string
	GroupChats []*GroupChat `gorm:"many2many:user_group_chat;"`
}
