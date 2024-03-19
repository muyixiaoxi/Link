package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupRemarkLogic {
	return &UpdateGroupRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupRemarkLogic) UpdateGroupRemark(in *user.UpdateGroupRemarkRequest) (*user.UpdateGroupRemarkResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateGroupRemarkResponse{}, nil
}
