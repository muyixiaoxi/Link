package logic

import (
	"context"
	"fmt"
	"user/service/internal/types"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserQueryPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserQueryPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserQueryPhoneLogic {
	return &UserQueryPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserQueryPhoneLogic) UserQueryPhone(in *user.UserQueryPhoneRequest) (resp *user.UserQueryPhoneResponse, err error) {
	fmt.Println(in.Phone)
	model := &types.User{}
	err = l.svcCtx.DB.Where("phone = ?", in.Phone).First(model).Error
	if err != nil {
		return
	}
	resp = &user.UserQueryPhoneResponse{
		Id:        uint64(model.ID),
		Username:  model.Username,
		Avatar:    model.Avatar,
		Address:   model.Address,
		Signature: model.Signature,
	}
	l.svcCtx.DB.Table("tb_tag").Select("IFNULL(tag_name,group_name) tag_name").Joins("join tb_user_tag on tb_tag.id = tb_user_tag.tag_id").Where("tb_user_tag.user_id = ?", resp.Id).Scan(&resp.TagName)
	return
}
