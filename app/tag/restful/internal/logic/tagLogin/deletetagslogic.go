package tagLogin

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"tag/service/tag"
)

type DeleteTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagsLogic {
	return &DeleteTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagsLogic) DeleteTags(req *types.DelteTags) error {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("user_id").(json.Number)
	id, _ := userId.Int64()
	//封装请求参数
	deleteTagsRpcParams := &tag.DeleteTagRequest{
		CreatorId: uint64(id),
		TagId:     req.TagIds,
	}
	_, err := l.svcCtx.TagLogin.DeleteTag(l.ctx, deleteTagsRpcParams)
	if err != nil {
		return err
	}
	return nil
}
