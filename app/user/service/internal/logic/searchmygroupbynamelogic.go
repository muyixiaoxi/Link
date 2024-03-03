package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMyGroupByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchMyGroupByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMyGroupByNameLogic {
	return &SearchMyGroupByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchMyGroupByNameLogic) SearchMyGroupByName(in *user.SearchMyGroupByNameRequest) (*user.MyGroupResponse, error) {
	// 根据备注或者群名
	var groupId []uint64
	err := l.svcCtx.DB.Model(&types.UserGroupChat{}).Select("group_chat_id").Where("user_id = ?", in.UserId).Find(&groupId).Error
	if err != nil {
		return nil, err
	}
	// 先根据备注或者搜索出我加入的群聊的id
	var myGroupID []uint64
	err = l.svcCtx.DB.Model(&types.UserGroupChat{}).Select("group_chat_id").Where("user_id = ? AND remark = ?", in.UserId, in.RemarkOrName+"%").Find(&myGroupID).Error
	if err != nil {
		return nil, err
	}
	var groupChat []types.GroupChat
	err = l.svcCtx.DB.Table("group_chats").Where("name LIKE ? AND id IN (?)", in.RemarkOrName+"%", groupId).Find(&groupChat).Error
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
		if err := l.svcCtx.DB.Table("tb_tag").Select("tag_name").Where("id = ?", groupInfo.UserSelfTagId).Scan(&tagName).Error; err != nil {
			return nil, err
		}
		var remark string
		if err := l.svcCtx.DB.Table("user_group_chats").Select("remark").Where("user_id = ? AND group_chat_id = ? AND remark IS NOT NULL", in.UserId, groupInfo.ID).Scan(&remark).Error; err != nil {
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
	resp := &user.MyGroupResponse{
		GroupList: groupInformations,
	}
	return resp, nil
}
