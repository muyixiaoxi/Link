package logic

import (
	"context"
	"encoding/json"
	"user/service/internal/types/model"

	"user/service/internal/svc"
	"user/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// ChatService 聊天服务
type ChatService struct {
	Hub *model.Hub //聊天服务
}

// NewChatService 函数创建一个新的ChatService实例
func NewChatService(hub *model.Hub) *ChatService {
	return &ChatService{
		Hub: hub,
	}
}

func NewGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupChatLogic {
	return &GroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupChatLogic) GroupChat(in *user.GroupChatRequest) (*user.Empty, error) {
	// 聊天的逻辑
	msg := model.Message{
		Content: in.Content,
		Type:    in.Type,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	l.svcCtx.RDB2.Publish(l.ctx, "chat", string(data)) //发布消息到redis中
	return &user.Empty{}, nil
}
