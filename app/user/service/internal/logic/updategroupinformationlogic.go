package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupInformationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupInformationLogic {
	return &UpdateGroupInformationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupInformationLogic) UpdateGroupInformation(in *user.UpdateGroupInfoRequest) (*user.UpdateGroupInfoResponse, error) {
	//判断当前用户是否是群主
	var groupBossId uint64
	err := l.svcCtx.DB.Model(&types.GroupChat{}).Where("id = ?", in.Id).Select("group_boss_id").Find(&groupBossId).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(groupBossId, in.UserId)
	if groupBossId != in.UserId {
		//不是群主
		return nil, status.Error(899, "无权操作")
	}
	// 修改群聊信息
	newInfo := &types.GroupChat{
		SystemTagId:   in.SystemTagId,
		UserSelfTagId: in.UserSelfTagId,
		Avatar:        in.Avatar,
		Name:          in.Name,
	}
	err = l.svcCtx.DB.Where("id = ?", in.Id).Updates(newInfo).Error
	if err != nil {
		return nil, err
	}
	return &user.UpdateGroupInfoResponse{}, nil
}
