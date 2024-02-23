package types

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserID   uint64 `gorm:"not null"`
	FriendID uint64 `gorm:"not null"`
	IsFriend bool   `gorm:"not null"`
	Remark   string `gorm:"size:20;"` // 好友备注
}
