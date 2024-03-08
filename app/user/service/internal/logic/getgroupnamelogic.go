package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupNameLogic {
	return &GetGroupNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupNameLogic) GetGroupName(in *user.GetGroupNameRequest) (*user.GetGroupNameResponse, error) {
	// 获取群聊名称
	var groupName string
	err := l.svcCtx.DB.Model(&types.GroupChat{}).Select("name").Where("id = ?", in.Id).Find(&groupName).Error
	return &user.GetGroupNameResponse{GroupName: groupName}, err
}
