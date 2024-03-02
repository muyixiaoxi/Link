package tagloginlogic

import (
	"context"
	"tag/service/internal/types"

	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelUserTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelUserTagLogic {
	return &CancelUserTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelUserTagLogic) CancelUserTag(in *tag.CancelRequest) (*tag.Empty, error) {
	// 用户取消选择标签
	err := l.svcCtx.DB.Where("user_id = ? and tag_id IN (?)", in.UserId, in.TagIds).Delete(&types.UserTagFollow{}).Error
	if err != nil {
		return nil, err
	}
	return &tag.Empty{}, nil
}
