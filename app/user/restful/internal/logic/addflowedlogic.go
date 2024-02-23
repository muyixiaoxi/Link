package logic

import (
	"context"
	"encoding/json"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFlowedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlowedLogic {
	return &AddFlowedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFlowedLogic) AddFlowed(req *types.UserAppleRequest) (err error) {
	jid := l.ctx.Value("user_id").(json.Number)
	id, _ := jid.Int64()
	req.From = uint64(id)
	_, err = l.svcCtx.UserRpc.UserFlowed(l.ctx, &user.UserAddRequest{
		Id:      uint64(id),
		BeId:    req.To,
		Message: req.Content,
		Type:    uint32(req.Type),
		Remark:  req.Remark,
	})
	if err != nil {
		logx.Error("l.svcCtx.UserRpc.UserFlowed failed: ", err)
		return
	}

	// 如果该用户登录
	if client, has := Clients[req.To]; has {
		client.Conn.WriteJSON(req)
	}
	return
	return nil
}
