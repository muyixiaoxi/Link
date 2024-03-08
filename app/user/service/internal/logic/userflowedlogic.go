package logic

import (
	"context"
	"errors"
	"user/service/internal/svc"
	"user/service/internal/types"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFlowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFlowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFlowedLogic {
	return &UserFlowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFlowedLogic) UserFlowed(in *user.UserAddRequest) (response *user.Empty, err error) {
	response = &user.Empty{}
	model := types.ApplyFor{
		UserID:  in.Id,
		BeId:    in.BeId,
		Message: in.Message,
		Type:    in.Type,
		Remark:  in.Remark,
	}
	// 首先判断是否为自己好友
	if model.Type == 3 {
		tmp := &types.Friend{}
		err = l.svcCtx.DB.Where("user_id = ? and friend_id = ?", model.UserID, model.BeId).First(tmp).Error
	} else {
		tmp := &types.UserGroupChat{}
		err = l.svcCtx.DB.Where("user_id = ? and group_chat_id = ?", model.UserID, model.BeId).First(tmp).Error
	}
	// 如果添加过该好友/群聊  返回
	if err == nil {
		return response, errors.New("repeat addition")
	}

	// 未添加过好友，如果是群聊，判断群聊人数
	if model.Type == 4 {
		var count int64
		l.svcCtx.DB.Model(&types.UserGroupChat{}).Where("group_chat_id = ?", model.BeId).Count(&count)
		tmp := &types.GroupChat{}
		tmp.ID = uint(model.BeId)
		l.svcCtx.DB.Find(tmp)
		if count >= int64(tmp.Max) {
			return response, errors.New("group chat overload")
		}
	}
	err = nil
	tmp := types.ApplyFor{}
	err = l.svcCtx.DB.Where("user_id = ? and be_id = ?", model.UserID, model.BeId).First(&tmp).Error
	if err != nil {
		err = l.svcCtx.DB.Create(&model).Error
		return
	}
	err = l.svcCtx.DB.Where("user_id = ? and be_id = ?", model.UserID, model.BeId).Updates(&model).Update("result", nil).Error
	return response, err
}
