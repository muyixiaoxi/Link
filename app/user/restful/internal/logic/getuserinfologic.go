package logic

import (
	"context"
	"user/common/jwt"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	userID := l.ctx.Value(jwt.UserId).(uint64)
	response, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user.UserInfoRequest{
		UserId:   userID,
		FriendId: req.Id,
	})
	if err != nil {
		logx.Error("l.svcCtx.UserRpc.UserInfo failed: ", err)
		return nil, err
	}
	resp = &types.UserInfoResponse{
		Id:        uint(response.Id),
		Username:  response.Username,
		Avatar:    response.Avatar,
		Age:       uint(response.Age),
		Gender:    uint(response.Gender),
		Address:   response.Address,
		Phone:     response.Phone,
		Signature: response.Signature,
		IsFriend:  response.IsFriend,
		Remark:    response.Remark,
	}
	return
}
