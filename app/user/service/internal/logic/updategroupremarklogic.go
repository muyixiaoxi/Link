package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupRemarkLogic {
	return &UpdateGroupRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupRemarkLogic) UpdateGroupRemark(in *user.UpdateGroupRemarkRequest) (*user.UpdateGroupRemarkResponse, error) {
	// 修改对群聊的备注
	remark := &types.UserGroupChat{
		Remark: in.Remark,
	}
	err := l.svcCtx.DB.Where("group_chat_id = ? and user_id = ?", in.Id, in.UserId).Updates(&remark).Error
	if err != nil {
		return nil, err
	}
	return &user.UpdateGroupRemarkResponse{}, nil
}
