package tagLogin

import (
	"context"
	"encoding/json"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"tag/service/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSelfTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSelfTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSelfTagLogic {
	return &UpdateSelfTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSelfTagLogic) UpdateSelfTag(req *types.UpdateSelfTagRequest) error {
	// 修改用户自定义标签
	userId := l.ctx.Value("user_id").(json.Number)
	id, _ := userId.Int64()
	updateSelfTagParam := &tag.CreateTagRequest{
		CreatorId:  uint64(id),  //当前登录的用户
		TagName:    req.TagName, //新的标签名称
		GroupName:  req.GroupName,
		OldTagName: req.OldName, //旧的标签名称
	}
	_, err := l.svcCtx.TagLogin.UpdateTag(l.ctx, updateSelfTagParam)
	if err != nil {
		return err
	}
	return nil
}
