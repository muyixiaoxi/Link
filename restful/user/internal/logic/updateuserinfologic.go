package logic

import (
	"Link/service/user/user"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"

	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"

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
	jid := l.ctx.Value("user_id").(json.Number)
	id, _ := jid.Int64()
	_, err := l.svcCtx.UserRpc.UserUpdateInfo(l.ctx, &user.UserUpdateInfoRequest{
		Id:       uint64(id),
		Username: req.Username,
		Password: req.Password,
		Avatar:   req.Avatar,
		Age:      uint64(req.Age),
		Gender:   uint64(req.Gender),
		Address:  req.Address,
		Phone:    req.Phone,
	})
	if err != nil {
		logc.Error(context.Background(), "l.svcCtx.UserRpc.UserUpdateInfo failed: ", err)
		return err
	}
	return nil
}
