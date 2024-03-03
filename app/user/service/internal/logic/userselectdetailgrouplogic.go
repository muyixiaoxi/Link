package logic

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSelectDetailGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// AddressCount 统计人数最多的地址
type AddressCount struct {
	Address string
	Count   int64
}

func NewUserSelectDetailGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSelectDetailGroupLogic {
	return &UserSelectDetailGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSelectDetailGroupLogic) UserSelectDetailGroup(in *user.DetailGroupRequest) (*user.DetailGroupResponse, error) {
	// 查询群聊详细信息
	var groupInfo types.GroupChat
	if err := l.svcCtx.DB.Where("id = ?", in.Id).Find(&groupInfo).Error; err != nil {
		return nil, err
	}
	// 查询该群聊所属的标签
	var (
		groupName   string
		tagName     string
		totalPeople int64
	)
	if err := l.svcCtx.DB.Table("tb_tag").Select("group_name").Where("id = ?", groupInfo.SystemTagId).Find(&groupName).Error; err != nil {
		return nil, err
	}
	if err := l.svcCtx.DB.Table("tb_tag").Select("tag_name").Where("id = ?", groupInfo.UserSelfTagId).Find(&tagName).Error; err != nil {
		return nil, err
	}
	//统计群聊中的人数
	var userIds []int64
	if err := l.svcCtx.DB.Model(&types.UserGroupChat{}).Select("user_id").Where("group_chat_id = ?", groupInfo.ID).Find(&userIds).Error; err != nil {
		return nil, err
	}
	if err := l.svcCtx.DB.Model(&types.UserGroupChat{}).Where("group_chat_id = ?", groupInfo.ID).Count(&totalPeople).Error; err != nil {
		return nil, err
	}
	//统计群聊中男女比列 以及 人数最多的地区
	var (
		manCount   int64
		womanCount int64
	)
	if err := l.svcCtx.DB.Model(&types.User{}).Where("id in (?) AND gender = 1", userIds).Count(&manCount).Error; err != nil {
		return nil, err
	}
	if err := l.svcCtx.DB.Model(&types.User{}).Where("id in (?) AND gender = 2", userIds).Count(&womanCount).Error; err != nil {
		return nil, err
	}
	//统计人数最多的地址
	var mostPopularAddress AddressCount
	if err := l.svcCtx.DB.Model(&types.User{}).
		Select("address, count(*) as count").
		Where("id IN (?) AND address IS NOT NULL", userIds).
		Group("address").
		Order("count DESC").
		Limit(1).
		Find(&mostPopularAddress).
		Error; err != nil {
		return nil, err
	}
	// 创建一个 *timestamppb.Timestamp 对象
	timestampProto := timestamppb.New(groupInfo.CreatedAt)
	return &user.DetailGroupResponse{
		Id:              uint64(groupInfo.ID),
		Count:           uint64(totalPeople),
		GroupBossId:     groupInfo.GroupBossID,
		Name:            groupInfo.Name,
		SystemTagName:   groupName,
		UserSelfTagName: tagName,
		Avatar:          groupInfo.Avatar,
		CreatedAt:       timestampProto,
		Man:             float32(manCount) / float32(totalPeople),
		Woman:           float32(womanCount) / float32(totalPeople),
		Address:         mostPopularAddress.Address,
		AddressCount:    float32(mostPopularAddress.Count),
	}, nil
}
