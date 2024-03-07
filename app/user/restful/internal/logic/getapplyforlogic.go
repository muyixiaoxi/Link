package logic

import (
	"context"
	"encoding/json"
	"time"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyForLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplyForLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyForLogic {
	return &GetApplyForLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplyForLogic) GetApplyFor() (resp []types.UserApplyForResponse, err error) {
	jid := l.ctx.Value("user_id").(json.Number)
	id, _ := jid.Int64()
	response, err := l.svcCtx.UserRpc.UserGetApplyFor(context.Background(), &user.UserGetApplyForRequest{
		UserId: uint64(id),
	})
	resp = make([]types.UserApplyForResponse, len(response.List))
	for i, applyFor := range response.List {
		t, _ := time.Parse(time.RFC3339Nano, applyFor.UpdatedAt)
		time := t.Format("2006/01/02 15:04:05")
		tmp := types.UserApplyForResponse{
			UserId:    applyFor.UserId,
			Username:  applyFor.Username,
			Avatar:    applyFor.Avatar,
			BeId:      applyFor.BeId,
			Message:   applyFor.Message,
			Type:      applyFor.Type,
			Result:    applyFor.Result,
			UpdatedAt: time,
		}
		resp[i] = tmp
	}
	return
}
