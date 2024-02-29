package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPhoneLogic {
	return &QueryPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPhoneLogic) QueryPhone(req *types.UserQueryPhoneRequest) (resp *types.UserQueryPhoneResponse, err error) {
	response, err := l.svcCtx.UserRpc.UserQueryPhone(context.Background(), &user.UserQueryPhoneRequest{Phone: req.Phone})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserQueryPhone failed: ", err)
		return
	}
	resp = &types.UserQueryPhoneResponse{
		Id:        response.Id,
		Username:  response.Username,
		Avatar:    response.Avatar,
		Address:   response.Address,
		Signature: response.Signature,
		TagName:   response.TagName,
	}
	return
}
