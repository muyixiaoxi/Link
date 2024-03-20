package types

type Message struct {
	ID          string `gorm:"primary_key"`
	From        uint64 `gorm:"not null"`
	To          uint64 `gorm:"not null"`
	Type        uint8  `gorm:"not null"`
	ContentType uint8  `gorm:"not null;comment:0 文字 1 图片"`
	Time        string `gorm:"not null"`
	Content     string `gorm:"type:varchar(500)"`
	Read        bool   `gorm:"default:0;comment:已读"`
}
