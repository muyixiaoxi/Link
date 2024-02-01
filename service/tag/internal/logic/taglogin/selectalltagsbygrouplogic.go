package tagloginlogic

import (
	"context"

	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllTagsByGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectAllTagsByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllTagsByGroupLogic {
	return &SelectAllTagsByGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectAllTagsByGroupLogic) SelectAllTagsByGroup(in *tag.SelectAllTagsByGroupName) (*tag.AllTagsByGroupNameResponse, error) {
	// todo: add your logic here and delete this line

	return &tag.AllTagsByGroupNameResponse{}, nil
}