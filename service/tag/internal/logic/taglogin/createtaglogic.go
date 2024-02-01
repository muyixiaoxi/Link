package tagloginlogic

import (
	"Link/service/tag/internal/types"
	"context"
	"errors"

	"Link/service/tag/internal/svc"
	"Link/service/tag/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateTag 用户自定义标签
func (l *CreateTagLogic) CreateTag(in *tag.CreateTagRequest) (*tag.CreateTagResponse, error) {
	// todo: add your logic here and delete this line
	var createRequst types.Tag
	err := l.svcCtx.DB.Take(&createRequst, "tag_name = ?", in.TagName).Error
	if err == nil {
		return nil, errors.New("TAG IS EXISTS")
	}
	//如果标签不存在则创建
	createRequst = types.Tag{
		Type:      "OFFICIAL",
		GroupName: "游戏",
		TagName:   "",
		CreatorID: -1,
	}
	err = l.svcCtx.DB.Create(&createRequst).Error
	//根据groupName 查询小标签

	return &tag.CreateTagResponse{}, err
}
