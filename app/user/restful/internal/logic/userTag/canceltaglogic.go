package userTag

import (
	"context"
	"encoding/json"
	"user/service/tag/service/tag"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTagLogic {
	return &CancelTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelTagLogic) CancelTag(req *types.CancelRequest) error {
	// 取消选择标签
	jid := l.ctx.Value("user_id").(json.Number)
	userId, _ := jid.Int64()

	_, err := l.svcCtx.TagLoginRpc.CancelUserTag(l.ctx, &tag.CancelRequest{
		UserId: uint64(userId),
		TagIds: req.TagIds,
	})
	return err
}
