package types

import (
	"gorm.io/gorm"
	"time"
)

// ChatHistory 聊天记录表
type ChatHistory struct {
	gorm.Model
	FromUserID    uint64    `json:"from_user_id" gorm:"comment:'消息发送者的ID'"`
	ToID          uint64    `json:"to_user_id" gorm:"comment:'消息接收的人或者群聊的ID'"`
	Type          uint8     `json:"type" gorm:"comment:'消息类型 1:单聊 2:群聊'"`
	MessageIsRead uint8     `json:"message_is_read" gorm:"comment:'消息已读未读 1:已读 2:未读'"`
	Time          time.Time `json:"time" gorm:"comment:'消息发送时间'"`
	Content       string    `json:"content" gorm:"comment:'消息内容'"`
}

func (c *ChatHistory) TableName() string {
	return "tb_chat_history"
}
