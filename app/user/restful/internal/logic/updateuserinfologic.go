package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"user/common/jwt"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UserUpdateInfoRequest) error {
	userID := l.ctx.Value(jwt.UserId).(uint64)
	_, err := l.svcCtx.UserRpc.UserUpdateInfo(l.ctx, &user.UserUpdateInfoRequest{
		Id:        userID,
		Username:  req.Username,
		Password:  req.Password,
		Avatar:    req.Avatar,
		Age:       uint64(req.Age),
		Gender:    uint64(req.Gender),
		Address:   req.Address,
		Phone:     req.Phone,
		Signature: req.Signature,
	})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserUpdateInfo failed: ", err)
	}
	return err
}
