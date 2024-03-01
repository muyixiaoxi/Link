package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSelectGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSelectGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectGroupLogic {
	return &UserSelectGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectGroupLogic) UserSelectGroup(in *user.UserSelectGroupsRequest) (resp *user.UserSelectGroupsResponse, err error) {
	// 根据标签查询群聊 首先根据小标签推选群聊 , 如果小标签不存在,则根据大标签推荐
	var groupInformations []*user.GroupInformation
	var groupChat []types.GroupChat
	var total int64
	db := l.svcCtx.DB
	page := int(in.PageNo)
	pageSize := int(in.PageSize)
	offset := (page - 1) * pageSize
	query := db.Model(&types.GroupChat{})
	//把用户选择的大标签和小标签的数据全部查询出来
	query = query.Where("system_tag_id in (?) or user_self_tag_id in (?)", in.TagId, in.TagId)
	// 查询当前页的数据
	err = query.Offset(offset).Limit(pageSize).Find(&groupChat).Error
	if err != nil {
		return nil, err
	}
	// 查询总记录数
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	//根据标签ID查询系统标签名称并添加到结果中
	for _, groupInfo := range groupChat {
		var systemTagName string
		if err := db.Table("tb_tag").Select("group_name").Where("id = ? ", groupInfo.SystemTagId).Scan(&systemTagName).Error; err != nil {
			return nil, err
		}
		var tagName string
		if err := db.Table("tb_tag").Select("tag_name").Where("id = ?", groupInfo.UserSelfTagId).Scan(&tagName).Error; err != nil {
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
	resp = &user.UserSelectGroupsResponse{
		GroupList: groupInformations,
		Total:     uint64(total),
	}

	return
}
