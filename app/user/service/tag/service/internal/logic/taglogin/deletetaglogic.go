package tagloginlogic

import (
	"context"
	"user/service/tag/service/internal/svc"
	"user/service/tag/service/internal/types"
	"user/service/tag/service/tag"

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
