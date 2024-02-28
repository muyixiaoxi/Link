package tagsignlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUserChooseTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUserChooseTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUserChooseTagLogic {
	return &SignUserChooseTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUserChooseTagLogic) SignUserChooseTag(in *tag.UserChooseTagRequest) (*tag.UserChooseTagRequest, error) {
	// todo: add your logic here and delete this line

	return &tag.UserChooseTagRequest{}, nil
}
