package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateRemarkLogic {
	return &UserUpdateRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateRemarkLogic) UserUpdateRemark(in *user.UserUpdateRemarkRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line

	return &user.Empty{}, nil
}
