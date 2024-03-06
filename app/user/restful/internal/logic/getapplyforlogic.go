package logic

import (
	"context"
	"encoding/json"
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
	for _, applyFor := range response.List {
		tmp := types.UserApplyForResponse{
			UserId:    applyFor.UserId,
			BeId:      applyFor.BeId,
			Message:   applyFor.Message,
			Type:      applyFor.Type,
			Result:    applyFor.Result,
			UpdatedAt: applyFor.UpdatedAt,
		}
		resp = append(resp, tmp)
	}
	return
}
