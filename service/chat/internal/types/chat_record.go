package types

import "gorm.io/gorm"

type ChatRecord struct {
	gorm.Model
	UserId uint64 `gorm:"not null"`
	BeId   uint64 `gorm:"not null"`
	Record string `gorm:"size:10;not null"`
	IsRead uint32 `gorm:"not null"`
	Type   uint32 `gorm:"not null"`
}
