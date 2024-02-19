package types

import "gorm.io/gorm"

type Tag struct {
	*gorm.Model
	Type      string `json:"type"`       //标签类型
	GroupName string `json:"group_name"` //标签所属组织类型
	TagName   string `json:"tag_name"`   //标签名称
	CreatorID int64  `json:"creator_id"` //创建这个标签的用户id
}

func (t *Tag) TableName() string {
	return "tb_tag"
}
