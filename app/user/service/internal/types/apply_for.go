package types

import "gorm.io/gorm"

type ApplyFor struct {
	gorm.Model
	UserID  uint64 `gorm:"not null"`
	BeId    uint64 `gorm:"not null"`
	Message string `gorm:"size:100"`
	Type    uint32 `gorm:"not null"`
}
