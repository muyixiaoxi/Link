package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
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
	if req.Id == 0 {
		jId := l.ctx.Value("user_id").(json.Number)
		id, _ := jId.Int64()
		req.Id = uint(id)
	}
	response, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uint64(req.Id),
	})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserInfo failed: ", err)
		return
	}
	resp = &types.UserInfoResponse{
		Id:       uint(response.Id),
		Username: response.Username,
		Avatar:   response.Avatar,
		Age:      uint(response.Age),
		Gender:   response.Gender,
		Address:  response.Address,
		Phone:    response.Phone,
	}
	return
}
