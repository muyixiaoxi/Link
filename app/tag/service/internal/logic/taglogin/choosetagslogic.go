package tagloginlogic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tag/service/internal/types"

	"tag/service/internal/svc"
	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChooseTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChooseTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChooseTagsLogic {
	return &ChooseTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChooseTagsLogic) ChooseTags(in *tag.ChooseTagsRequest) (resp *tag.ChooseTagsResponse, err error) {
	//判断大标签是否被选择过
	tx := l.svcCtx.DB.Begin()
	err = tx.Take(&types.UserTagFollow{}, "tag_id = ? and user_id = ?", in.SystemId, in.Id).Error
	if err != nil {
		//大标签没有被选择,插入
		err = tx.Create(&types.UserTagFollow{
			TagId:  in.SystemId,
			UserId: in.Id,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// 插入用户选择的小标签
	for _, tagId := range in.TagIds {
		var userTag types.UserTagFollow
		// 尝试查找记录
		err = tx.Where("tag_id = ? AND user_id = ?", tagId, in.Id).First(&userTag).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果记录不存在，则插入新记录
			err = tx.Create(&types.UserTagFollow{
				TagId:  tagId,
				UserId: in.Id,
			}).Error
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	return &tag.ChooseTagsResponse{}, tx.Commit().Error
}
