package logic

import (
	"context"
	"gorm.io/gorm"
	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecommendUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendUsersLogic {
	return &RecommendUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecommendUsersLogic) RecommendUsers(in *user.RecommendUsersRequest) (resp *user.RecommendUsersResponse, err error) {
	//推荐人员信息
	//先查询出符合条件的人员总数
	resp, err = QueryUsersByConditions(l.svcCtx.DB, in.PageNo, in.PageSize, in.TagId, in.UserId)
	return resp, err
}

func QueryUsersByConditions(db *gorm.DB, pageNo, pageSize int64, tagIds []uint64, userID uint64) (*user.RecommendUsersResponse, error) {
	var totalCount int64
	var userList []*user.RecommendUser
	offset := (pageNo - 1) * pageSize
	if err := db.
		Table("users").
		Joins("JOIN tb_user_tag ON users.id = tb_user_tag.user_id AND users.deleted_at IS NULL AND tb_user_tag.deleted_at IS NULL").
		Joins("JOIN tb_tag ON tb_user_tag.tag_id = tb_tag.id").
		Where("tb_tag.id IN (?) AND users.id != ? AND users.deleted_at IS NULL", tagIds, userID).
		Group("users.id").
		Count(&totalCount).Error; err != nil {
		return nil, err
	}
	if err := db.Debug().
		Table("users").
		Joins("JOIN tb_user_tag ON users.id = tb_user_tag.user_id AND users.deleted_at IS NULL AND tb_user_tag.deleted_at IS NULL").
		Joins("JOIN tb_tag ON tb_user_tag.tag_id = tb_tag.id").
		Where("tb_tag.id IN (?) AND users.id != ? AND users.deleted_at IS NULL", tagIds, userID).
		Select("users.username, users.avatar, users.signature, users.id").
		Group("users.id").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&userList).Error; err != nil {
		return nil, err
	}
	return &user.RecommendUsersResponse{RecommendUserList: userList, Total: totalCount}, nil
}
