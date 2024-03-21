package types

type GroupMessage struct {
	ID          string `json:"ID" gorm:"primary_key"`
	From        uint64 `json:"from"`
	To          uint64 `json:"to"`
	Type        uint8  `json:"type"`
	ContentType uint8  `json:"content_type"`
	Time        string `json:"time"`
	Content     string `json:"content"`
	Read        bool   `json:"read" gorm:"default:0"`
}
