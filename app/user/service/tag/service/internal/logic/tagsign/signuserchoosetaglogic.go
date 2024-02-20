package tagsignlogic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tag/service/internal/types"

	"tag/service/internal/svc"
	"tag/service/tag"

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
	var chooseTag types.UserTagFollow
	// 用户注册时选择标签
	err := l.svcCtx.DB.Take(&chooseTag, "tag_id = ? and user_id = ? and type = 'OFFICIAL'", in.TagId, in.UserId).Error
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "禁止重复选择标签")
	}
	chooseTag = types.UserTagFollow{
		TagId:  in.TagId,
		UserId: in.UserId,
	}
	err = l.svcCtx.DB.Create(&chooseTag).Error
	return &tag.UserChooseTagRequest{}, nil
}
