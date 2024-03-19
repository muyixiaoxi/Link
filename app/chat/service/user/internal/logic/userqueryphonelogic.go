package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserQueryPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryPhoneLogic {
	return &UserQueryPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserQueryPhoneLogic) UserQueryPhone(in *user.UserQueryPhoneRequest) (*user.UserQueryPhoneResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserQueryPhoneResponse{}, nil
}
