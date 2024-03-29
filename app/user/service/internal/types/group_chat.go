package types

import "gorm.io/gorm"

type GroupChat struct {
	gorm.Model
	GroupBossID   uint64 `gorm:"not null" json:"group_boss_id"`
	SystemTagId   uint64 `json:"system_tag_id"`
	UserSelfTagId uint64 `json:"user_self_tag_id"`
	Avatar        string `json:"avatar"`
	Name          string `gorm:"size:20;not null" json:"name"`
	Max           uint32 `gorm:"default:50"`
}
