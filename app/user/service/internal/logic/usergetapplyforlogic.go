package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetApplyForLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGetApplyForLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetApplyForLogic {
	return &UserGetApplyForLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGetApplyForLogic) UserGetApplyFor(in *user.UserGetApplyForRequest) (resp *user.UserGetApplyForResponse, err error) {
	resp = &user.UserGetApplyForResponse{}
	err = l.svcCtx.DB.Model(&types.ApplyFor{}).Where("be_id = ?", in.UserId).Scan(&resp.List).Error
	return
}
