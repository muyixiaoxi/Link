package tagLogin

import (
	"context"
	"encoding/json"
	"tag/service/tag"

	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectMyTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectMyTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectMyTagsLogic {
	return &SelectMyTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectMyTagsLogic) SelectMyTags() (resp []types.SelectMyTagsResponse, err error) {
	// 查询我创建的标签
	userId := l.ctx.Value("user_id").(json.Number)
	id, _ := userId.Int64()
	myTagList, err := l.svcCtx.TagLogin.SelectMyTags(l.ctx, &tag.SelectMyTagsRequest{UserId: uint64(id)})
	if err != nil {
		return nil, err
	}
	for _, myTag := range myTagList.LowTags {
		temp := types.SelectMyTagsResponse{
			Id:      myTag.Id,
			TagName: myTag.TagName,
		}
		resp = append(resp, temp)
	}
	return
}
