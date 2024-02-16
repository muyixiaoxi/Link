package tagTest

import (
	"Link/service/tag/tag"
	"context"

	"Link/restful/tag/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTagLogic {
	return &UserTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTagLogic) UserTag() error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.TagTest.SignUserChooseTag(l.ctx, &tag.UserChooseTagRequest{
		UserId: 6,
		TagId:  1,
	})
	return err
}
