package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStrangerGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchStrangerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStrangerGroupLogic {
	return &SearchStrangerGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchStrangerGroupLogic) SearchStrangerGroup(in *user.SearchStrangerGroupRequest) (*user.UserSelectGroupsResponse, error) {
	// 搜索陌生群聊
	// 先搜索出我加入的群聊的id
	var myGroupID []uint64
	err := l.svcCtx.DB.Debug().Model(&types.UserGroupChat{}).Select("group_chat_id").Where("user_id = ?", in.UserId).Find(&myGroupID).Error
	if err != nil {
		return nil, err
	}
	//根据名称搜索我未加入的群聊
	page := int(in.PageNo)
	pageSize := int(in.PageSize)
	offset := (page - 1) * pageSize
	var groupChat []types.GroupChat
	err = l.svcCtx.DB.Table("group_chats").Where("name LIKE ? AND id NOT IN (?)", in.Name+"%", myGroupID).Offset(offset).Limit(pageSize).Find(&groupChat).Error
	if err != nil {
		return nil, err
	}
	//获取符合条件的总数
	var total int64
	err = l.svcCtx.DB.Table("group_chats").Where("name LIKE ? AND id NOT IN (?)", in.Name+"%", myGroupID).Count(&total).Error
	if err != nil {
		return nil, err
	}
	//根据标签ID查询系统标签名称并添加到结果中
	var groupInformations []*user.GroupInformation
	for _, groupInfo := range groupChat {
		var systemTagName string
		if err := l.svcCtx.DB.Table("tb_tag").Select("group_name").Where("id = ? ", groupInfo.SystemTagId).Scan(&systemTagName).Error; err != nil {
			return nil, err
		}
		var tagName string
		if err := l.svcCtx.DB.Table("tb_tag").Select("tag_name").Where("id = ? AND tag_name IS NOT NULL", groupInfo.UserSelfTagId).Scan(&tagName).Error; err != nil {
			return nil, err
		}
		temp := &user.GroupInformation{
			Id:              uint64(groupInfo.ID),
			Name:            groupInfo.Name,
			SystemTagName:   systemTagName,
			UserSelfTagName: tagName,
			Avatar:          groupInfo.Avatar,
		}
		groupInformations = append(groupInformations, temp)
	}
	resp := &user.UserSelectGroupsResponse{
		GroupList: groupInformations,
		Total:     uint64(total),
	}
	return resp, nil
}
