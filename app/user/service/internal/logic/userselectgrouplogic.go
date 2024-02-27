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
	var total int64
	db := l.svcCtx.DB
	page := int(in.PageNo)
	pageSize := int(in.PageSize)
	offset := (page - 1) * pageSize
	if in.UserSelfTagId == 0 {
		//只去查询和大标签相关的群聊 Limit每页显示的条数 Offset 当前页
		err = db.Model(&types.GroupChat{}).Where("system_tag_id = ?", in.SystemTagId).Limit(pageSize).Offset(offset).Scan(&groupInformations).Limit(-1).Offset(-1).Count(&total).Error
	} else {
		err = db.Model(&types.GroupChat{}).Where("user_self_tag_id = ?", in.UserSelfTagId).Limit(pageSize).Offset(offset).Scan(&groupInformations).Limit(-1).Offset(-1).Count(&total).Error
	}
	resp = &user.UserSelectGroupsResponse{
		GroupList: groupInformations,
		Total:     uint64(total),
	}
	return
}
