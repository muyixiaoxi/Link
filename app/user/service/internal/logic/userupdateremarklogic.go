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
	model := &types.Friend{
		UserID:   in.Id,
		FriendID: in.BeId,
		Remark:   in.Remark,
	}
	n := l.svcCtx.DB.Where("user_id = ? and friend_id = ?", in.Id, in.BeId).Updates(model).RowsAffected
	if n == 0 {
		err = l.svcCtx.DB.Create(model).Error
	}
	return
}
