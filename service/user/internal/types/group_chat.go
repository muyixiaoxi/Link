package types

import "gorm.io/gorm"

type GroupChat struct {
	gorm.Model
	UserID uint64  `gorm:"not null"`
	Name   string  `gorm:"size:20;not null"`
	Users  []*User `gorm:"many2many:user_group_chat;"`
}
