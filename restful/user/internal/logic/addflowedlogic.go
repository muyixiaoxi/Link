package logic

import (
	"Link/service/user/user"
	"context"
	"encoding/json"
	"fmt"

	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"

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

func (l *AddFlowedLogic) AddFlowed(req *types.UserFlowedRequest) (err error) {
	jid := l.ctx.Value("user_id").(json.Number)
	id, _ := jid.Int64()
	resp, err := l.svcCtx.UserRpc.UserFlowed(l.ctx, &user.UserAddRequest{
		Id:      uint64(id),
		BeId:    req.BeId,
		Message: req.Message,
		Type:    req.Type,
	})
	if err != nil {
		return
	}
	fmt.Println(resp)
	return
}
