package logic

import (
	"context"

	"chat/service/user/internal/svc"
	"chat/service/user/user"

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

func (l *RecommendUsersLogic) RecommendUsers(in *user.RecommendUsersRequest) (*user.RecommendUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &user.RecommendUsersResponse{}, nil
}
