package types

import "gorm.io/gorm"

type UserGroupChat struct {
	gorm.Model
	GroupChatID uint64 `json:"group_chat_id"`
	UserID      uint64 `json:"user_id"`
	Remark      string `json:"remark"`
}
