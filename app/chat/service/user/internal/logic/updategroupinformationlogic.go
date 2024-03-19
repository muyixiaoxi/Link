package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

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
	// todo: add your logic here and delete this line

	return &user.UpdateGroupInfoResponse{}, nil
}
