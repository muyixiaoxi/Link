package userGroup

import (
	"context"
	"user/common/jwt"
	"user/service/user"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupRemarkLogic {
	return &UpdateGroupRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupRemarkLogic) UpdateGroupRemark(req *types.UpdateGroupRemarkRequest) error {
	// 获取当前登录用户的id
	userID := l.ctx.Value(jwt.UserId).(uint64)
	// 修改群聊备注
	_, err := l.svcCtx.UserRpc.UpdateGroupRemark(l.ctx, &user.UpdateGroupRemarkRequest{
		Id:     req.GroupId,
		UserId: userID,
		Remark: req.Remark,
	})
	return err
}
