package tagloginlogic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tag/service/internal/types"

	"tag/service/internal/svc"
	"tag/service/tag"

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
	tx := l.svcCtx.DB.Begin()
	err := tx.Take(&createRequst, "tag_name = ?", in.TagName).Error
	if err == nil {
		tx.Rollback()
		return nil, status.Error(codes.AlreadyExists, "标签已经存在")
	}
	err = tx.Where("group_name = ?", in.GroupName).First(&createRequst).Error
	if err != nil {
		tx.Rollback()
		return nil, status.Error(codes.NotFound, "系统自定义标签不存在")
	}
	//如果标签不存在则创建
	createRequst = types.Tag{
		Type:      in.Type,
		GroupName: in.GroupName,
		TagName:   in.TagName,
		CreatorID: int64(in.CreatorId),
	}
	err = tx.Debug().Create(&createRequst).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	//根据groupName 查询小标签
	var lowTags []*tag.CreateTagResponse_LowTags
	err = l.svcCtx.DB.Where("group_name = ? and type != 'OFFICIAL'", in.GroupName).Model(&types.Tag{}).Find(&lowTags).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &tag.CreateTagResponse{
		GroupName: in.GroupName,
		LowTags:   lowTags,
	}, err
}
