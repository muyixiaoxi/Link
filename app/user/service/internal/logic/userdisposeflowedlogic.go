package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDisposeFlowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDisposeFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDisposeFlowedLogic {
	return &UserDisposeFlowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDisposeFlowedLogic) UserDisposeFlowed(in *user.DisposeFlowedRequest) (*user.Empty, error) {
	// 删除记录
	apply := &types.ApplyFor{
		UserID: in.From,
		BeId:   in.To,
		Type:   in.Type,
	}
	if err := l.svcCtx.DB.Where("user_id = ? and be_id = ? and type = ?", apply.UserID, apply.BeId, apply.Type).First(apply).Error; err != nil {
		return &user.Empty{}, err
	}
	tx := l.svcCtx.DB.Begin()
	if err := tx.Delete(apply).Error; err != nil {
		return &user.Empty{}, err
	}
	if in.Result { // 同意
		if in.Type == 3 { // 3 好友 4 群聊
			flowed1 := &types.Flowed{
				UserID:   in.From,
				FriendID: in.To,
			}
			flowed2 := &types.Flowed{
				UserID:   in.To,
				FriendID: in.From,
				Remark:   in.Remark,
			}
			if err := tx.Create(flowed1).Error; err != nil {
				tx.Rollback()
				return &user.Empty{}, err
			}
			if err := tx.Create(flowed2).Error; err != nil {
				tx.Rollback()
				return &user.Empty{}, err
			}
		} else if in.Type == 4 {
			userGroupChat := &types.UserGroupChat{
				UserID:      in.From,
				GroupChatID: in.To,
			}
			if err := tx.Create(userGroupChat).Error; err != nil {
				tx.Rollback()
				return &user.Empty{}, err
			}
		}
	}
	err := tx.Commit().Error

	return &user.Empty{}, err
}
