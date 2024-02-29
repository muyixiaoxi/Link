package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteFriendLogic {
	return &UserDeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteFriendLogic) UserDeleteFriend(in *user.UserDeleteFriendRequest) (resp *user.Empty, err error) {
	resp = &user.Empty{}
	m1 := &types.Friend{
		UserID:   in.UserId,
		FriendID: in.FriendId,
	}
	m2 := &types.Friend{
		UserID:   in.FriendId,
		FriendID: in.UserId,
	}
	tx := l.svcCtx.DB.Begin()
	err = tx.Delete(m1).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Delete(m2).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return resp, tx.Commit().Error
}
