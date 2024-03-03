package tagloginlogic

import (
	"context"
	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectLinkTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectLinkTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectLinkTagsLogic {
	return &SelectLinkTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectLinkTagsLogic) SelectLinkTags(in *tag.SelectLinkTagsRequest) (resp *tag.SelectLinkTagsResponse, err error) {
	// 查询当前登录用户创建或者加入的小标签
	// 如果是当前登录用户 , 只查询自己选择的标签 , 其它用户, 一起查询
	// 使用左外连接查询
	db := l.svcCtx.DB
	var linkTags []*tag.SelectLinkTag
	if in.IsSelf {
		//如果是当前登录用户,不能查询出自己创建的标签
		err = db.Table("tb_tag").Debug().
			Select("tb_tag.id, tb_tag.tag_name, tb_tag.creator_id , tb_tag.group_name").
			Joins("left join tb_user_tag on tb_tag.id = tb_user_tag.tag_id AND tb_tag.creator_id != ?", in.Id).
			Where("tb_user_tag.user_id = ? AND tb_user_tag.deleted_at IS NULL", in.Id).
			Where("tb_tag.deleted_at IS NULL").
			Find(&linkTags).Error
	} else {
		err = db.Table("tb_tag").
			Select("tb_tag.id, tb_tag.tag_name, tb_tag.creator_id , tb_tag.group_name").
			Joins("left join tb_user_tag on tb_tag.id = tb_user_tag.tag_id").
			Where("tb_user_tag.user_id = ? AND tb_user_tag.deleted_at IS NULL", in.Id).
			Where("tb_tag.deleted_at IS NULL").
			Find(&linkTags).Error
		if err != nil {
			return nil, err
		}
	}
	resp = &tag.SelectLinkTagsResponse{LinkTags: linkTags}
	return resp, err
}
