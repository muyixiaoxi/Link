package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

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

func (l *UserUpdateRemarkLogic) UserUpdateRemark(in *user.UserUpdateRemarkRequest) (response *user.Empty, err error) {
	response = &user.Empty{}
	model := types.Flowed{
		UserID:   in.Id,
		FriendID: in.BeId,
		Remark:   in.Remark,
	}
	err = l.svcCtx.DB.Updates(model).Error
	return
}
