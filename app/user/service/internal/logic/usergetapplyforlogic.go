package logic

import (
	"context"
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
	err = l.svcCtx.DB.Table("apply_fors a").Select("a.user_id,a.be_id,a.message,a.type,a.result,a.updated_at,u.username,u.avatar").Joins("join users u on u.id = a.user_id").Where("a.be_id = ?", in.UserId).Scan(&resp.List).Error
	return
}
