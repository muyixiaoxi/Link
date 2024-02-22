package types

import "gorm.io/gorm"

type Flowed struct {
	gorm.Model
	UserID   uint64 `gorm:"not null"`
	FriendID uint64 `gorm:"not null"`
	Remark   string `gorm:"size:20;"` // 好友备注
}