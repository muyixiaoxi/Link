package tagLogin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tag/common/jwt"
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
	userID := l.ctx.Value(jwt.UserId).(uint64)
	//封装请求参数
	deleteTagsRpcParams := &tag.DeleteTagRequest{
		CreatorId: userID,
		TagId:     req.TagIds,
	}
	_, err := l.svcCtx.TagLogin.DeleteTag(l.ctx, deleteTagsRpcParams)
	if err != nil {
		return err
	}
	return nil
}
