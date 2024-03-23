package logic

import (
	"chat/service/chat"
	"chat/service/internal/svc"
	"chat/service/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOfflineMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOfflineMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOfflineMessageLogic {
	return &SaveOfflineMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveOfflineMessage 存储离线消息
func (l *SaveOfflineMessageLogic) SaveOfflineMessage(in *chat.SaveMessageRequest) (response *chat.Empty, err error) {
	// 1. 存Mysql
	// 2. 存储到每个用户的redis
	// 3. redis存储成功
	// 4. 返回用户成功
	// 5. 用户下次登录时，修改Mysql中的消息状态
	message := &types.Message{
		ID:          in.Id,
		From:        in.From,
		To:          in.To,
		Type:        uint8(in.Type),
		ContentType: uint8(in.ContentType),
		Time:        string(in.Time),
		Content:     string(in.Content),
	}
	tx := l.svcCtx.DB.Begin()
	cache, _ := json.Marshal(message)
	err = tx.Save(message).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 如果离线需要存储 redis
	if !in.Online {
		err = l.SavaRedis(message, message.To)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	if err = tx.Commit().Error; err != nil {
		l.svcCtx.RDB.Zrem(fmt.Sprintf("link:chat:user:%d", in.To), string(cache))
	}
	return response, err
}

// SavaRedis redis 保存
func (l *SaveOfflineMessageLogic) SavaRedis(message *types.Message, userId uint64) (err error) {
	score, _ := strconv.Atoi(message.Time)
	cache, _ := json.Marshal(message)
	_, err = l.svcCtx.RDB.Zadd(fmt.Sprintf("link:chat:user:%d", userId), int64(score), string(cache))
	return
}
