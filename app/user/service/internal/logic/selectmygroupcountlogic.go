package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMyGroupCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectMyGroupCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMyGroupCountLogic {
	return &SelectMyGroupCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectMyGroupCountLogic) SelectMyGroupCount(in *user.SelectMyGroupCountRequest) (*user.SelectMyGroupCountResponse, error) {
	// 统计当前用户加入的群聊的个数
	var total int64
	err := l.svcCtx.DB.Debug().Model(&types.UserGroupChat{}).Where("user_id = ?", in.UserId).Count(&total).Error
	return &user.SelectMyGroupCountResponse{Count: uint64(total)}, err
}
