package logic

import (
	"context"
	"encoding/json"
	"user/restful/internal/svc"
	"user/restful/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRemarkLogic {
	return &UpdateRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRemarkLogic) UpdateRemark(req *types.UserUpdateRemarkRequest) (err error) {
	jId := l.ctx.Value("user_id").(json.Number)
	id, _ := jId.Int64()
	_, err = l.svcCtx.UserRpc.UserUpdateRemark(context.Background(), &user.UserUpdateRemarkRequest{
		Id:     uint64(id),
		BeId:   req.Friend,
		Remark: req.Remark,
	})
	return err
}
