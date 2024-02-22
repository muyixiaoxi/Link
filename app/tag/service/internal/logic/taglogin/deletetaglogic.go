package tagloginlogic

import (
	"context"
	"tag/service/internal/types"

	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTagLogic) DeleteTag(in *tag.DeleteTagRequest) (*tag.DeleteTagResponse, error) {
	// 删除用户自己创建的标签,同时,删除tb_user_tag标签中的记录
	tx := l.svcCtx.DB.Begin()
	err := tx.Where("creator_id = ? and id in (?) and type !='OFFICIAL'", in.CreatorId, in.TagId).Delete(&types.Tag{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//删除 tb_user_tag 标签中的记录
	var ids []int
	err = tx.Where("tag_id in (?)", in.TagId).Model(&types.UserTagFollow{}).Select("id").Find(&ids).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Where("id in (?)", ids).Delete(&types.UserTagFollow{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &tag.DeleteTagResponse{}, tx.Commit().Error
}
