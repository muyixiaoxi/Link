package tagLogin

import (
	"context"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"tag/service/tag"

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
	// 查询系统标签
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
