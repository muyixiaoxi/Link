package tagLogin

import (
	"Link/service/tag/tag"
	"context"
	"fmt"

	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectUserTagByGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectUserTagByGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectUserTagByGroupLogic {
	return &SelectUserTagByGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectUserTagByGroupLogic) SelectUserTagByGroup(req *types.SelectUserTagByGroupRequest) (resp *types.CreateTagResponse, err error) {
	// todo: add your logic here and delete this line
	respBody, err := l.svcCtx.TagLogin.SelectAllTagsByGroup(l.ctx, &tag.SelectAllTagsByGroupName{GroupAme: req.GroupName})
	fmt.Println(respBody)
	return
}
