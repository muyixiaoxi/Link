package logic

import (
	"context"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupHostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGroupHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGroupHostLogic {
	return &QueryGroupHostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryGroupHostLogic) QueryGroupHost(in *user.QueryGroupHostRequest) (resp *user.QueryGroupHostResponse, err error) {
	model := types.GroupChat{}
	err = l.svcCtx.DB.Where("id = ?", in.GroupId).Find(model).Error
	resp.GroupHostId = model.GroupBossID
	return resp, err
}
