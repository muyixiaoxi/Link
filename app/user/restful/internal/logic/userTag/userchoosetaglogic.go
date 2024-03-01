package userTag

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
	"user/service/tag/service/tag"

	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChooseTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserChooseTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChooseTagLogic {
	return &UserChooseTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserChooseTagLogic) UserChooseTag(req *types.UserChooseTagRequst) error {
	//获取当前登录的用户id
	jid := l.ctx.Value("user_id").(json.Number)
	id, _ := jid.Int64()
	//判断标签数量是否超过最大数量限制
	tagCount, err := l.svcCtx.TagLoginRpc.CheckTagCount(l.ctx, &tag.CheckTagCountRequest{UserId: uint64(id)})
	if tagCount.Count > 9 {
		return status.Error(899, "标签数量超过最大数量限制")
	}
	//选择标签
	_, err = l.svcCtx.TagLoginRpc.ChooseTags(l.ctx, &tag.ChooseTagsRequest{
		Id:       uint64(id),
		SystemId: req.SystemTagId,
		TagIds:   req.TagIds,
	})
	return err
}
