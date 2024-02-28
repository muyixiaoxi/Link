package tagsignlogic

import (
	"context"

	"user/service/tag/service/internal/svc"
	"user/service/tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUserChooseTagRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUserChooseTagRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUserChooseTagRevertLogic {
	return &SignUserChooseTagRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUserChooseTagRevertLogic) SignUserChooseTagRevert(in *tag.UserChooseTagRequest) (*tag.UserChooseTagRequest, error) {
	// todo: add your logic here and delete this line

	return &tag.UserChooseTagRequest{}, nil
}
