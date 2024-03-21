package logic

import (
	"chat/service/chat"
	"chat/service/internal/svc"
	"chat/service/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveGroupMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveGroupMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveGroupMessageLogic {
	return &SaveGroupMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveGroupMessageLogic) SaveGroupMessage(in *chat.SaveGroupMessageRequest) (*chat.Empty, error) {
	messageReq := in.Message
	message := &types.GroupMessage{
		ID:          messageReq.Id,
		From:        messageReq.From, //来自哪一个用户的
		To:          messageReq.To,   //发送到哪一个群的
		Type:        uint8(messageReq.Type),
		ContentType: uint8(messageReq.ContentType),
		Time:        string(messageReq.Time),
		Content:     messageReq.Content,
	}
	tx := l.svcCtx.DB.Begin() //mysql开启事务
	if err := tx.Save(message).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &chat.Empty{}, nil
}
