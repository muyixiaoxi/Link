package tagLogin

import (
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"Link/service/tag/tag"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectGroupTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectGroupTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectGroupTagLogic {
	return &SelectGroupTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectGroupTagLogic) SelectGroupTag() (resp []types.TagGroupName, err error) {
	// todo: add your logic here and delete this line
	empty := &tag.Empty{}
	respBody, err := l.svcCtx.TagLogin.SelectGroupTag(l.ctx, empty)
	if err != nil {
		return nil, err
	}
	for _, value := range respBody.Tags {
		temp := types.TagGroupName{
			GroupNameId: int(value.Id),
			GroupName:   value.GroupName,
		}
		resp = append(resp, temp)
	}
	return
}
