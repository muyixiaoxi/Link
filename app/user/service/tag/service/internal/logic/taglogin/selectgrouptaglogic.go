package tagloginlogic

import (
	"context"
	"user/service/tag/service/internal/svc"
	"user/service/tag/service/internal/types"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectGroupTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectGroupTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectGroupTagLogic {
	return &SelectGroupTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectGroupTagLogic) SelectGroupTag(in *tag.Empty) (*tag.GroupTagResponse, error) {
	// todo: add your logic here and delete this line
	var groupTag []*tag.GroupTag
	err := l.svcCtx.DB.Model(types.Tag{}).Where("type = 'OFFICIAL'").Find(&groupTag).Error
	return &tag.GroupTagResponse{
		Tags: groupTag,
	}, err
}
