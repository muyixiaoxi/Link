package tagloginlogic

import (
	"context"

	"Link/service/tag/internal/svc"
	"Link/service/tag/tag/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTagLogic) DeleteTag(in *tag.DeleteTagRequest) (*tag.DeleteTagResponse, error) {
	// todo: add your logic here and delete this line

	return &tag.DeleteTagResponse{}, nil
}
