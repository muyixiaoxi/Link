package tagloginlogic

import (
	"Link/service/tag/internal/types"
	"context"

	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"

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
	err := l.svcCtx.DB.Where("creator_id = ? and id in (?) and type !='OFFICIAL'", in.CreatorId, in.TagId).Delete(&types.Tag{}).Error
	if err != nil {
		return nil, err
	}
	return &tag.DeleteTagResponse{}, nil
}
