package logic

import (
	"context"
	"user/service/internal/svc"
	"user/service/internal/types"
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
	//查询当前登录用户创建的群id
	var groupIds []uint64
	err = l.svcCtx.DB.Model(&types.GroupChat{}).Select("id").Where("group_boss_id = ?", in.UserId).Find(&groupIds).Error
	if err != nil {
		return nil, err
	}
	groupIds = append(groupIds, in.UserId) //将当前用户创建的群聊的群聊申请以及好友申请全部查询出来
	err = l.svcCtx.DB.Table("apply_fors as a").Select("a.user_id,a.be_id,a.message,a.type,a.result,a.updated_at,u.username,u.avatar").Joins("join users u on u.id = a.user_id").Where("a.be_id in (?)", groupIds).Scan(&resp.List).Error
	return
}
