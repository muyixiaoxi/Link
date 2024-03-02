package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMyGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMyGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMyGroupListLogic {
	return &QueryMyGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMyGroupListLogic) QueryMyGroupList(in *user.QueryMyGroupListRequest) (*user.UserSelectGroupsResponse, error) {
	// 获取我加入的群聊列表
	page := int(in.PageNo)
	pageSize := int(in.PageSize)
	offset := (page - 1) * pageSize
	var groupChat []types.GroupChat
	err := l.svcCtx.DB.Table("group_chats").Joins("LEFT JOIN user_group_chats AS min ON min.group_chat_id = group_chats.id").Where("min.user_id = ?", in.UserId).Offset(offset).Limit(pageSize).Find(&groupChat).Error
	if err != nil {
		return nil, err
	}
	//获取符合条件的总数
	var total int64
	err = l.svcCtx.DB.Table("group_chats").Joins("LEFT JOIN user_group_chats AS min ON min.group_chat_id = group_chats.id").Where("min.user_id = ?", in.UserId).Count(&total).Error
	if err != nil {
		return nil, err
	}
	var groupInformations []*user.GroupInformation
	//根据标签ID查询系统标签名称并添加到结果中
	for _, groupInfo := range groupChat {
		var systemTagName string
		if err := l.svcCtx.DB.Table("tb_tag").Select("group_name").Where("id = ? ", groupInfo.SystemTagId).Scan(&systemTagName).Error; err != nil {
			return nil, err
		}
		var tagName string
		if err := l.svcCtx.DB.Table("tb_tag").Select("tag_name").Where("id = ?", groupInfo.UserSelfTagId).Scan(&tagName).Error; err != nil {
			return nil, err
		}
		var remark string
		if err := l.svcCtx.DB.Table("user_group_chats").Select("remark").Where("group_chat_id = ? and user_id = ? and remark IS NOT NULL", groupInfo.ID, in.UserId).Scan(&remark).Error; err != nil {
			return nil, err
		}
		temp := &user.GroupInformation{
			Id:              uint64(groupInfo.ID),
			Name:            groupInfo.Name,
			SystemTagName:   systemTagName,
			UserSelfTagName: tagName,
			Avatar:          groupInfo.Avatar,
			Remark:          remark,
		}
		groupInformations = append(groupInformations, temp)
	}
	resp := &user.UserSelectGroupsResponse{
		GroupList: groupInformations,
		Total:     uint64(total),
	}

	return resp, nil
}
