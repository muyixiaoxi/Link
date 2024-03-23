package logic

import (
	"chat/service/chat"
	"chat/service/internal/svc"
	"chat/service/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveGroupMessageRedisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveGroupMessageRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveGroupMessageRedisLogic {
	return &SaveGroupMessageRedisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveGroupMessageRedisLogic) SaveGroupMessageRedis(in *chat.SaveGroupMessageRequest) (*chat.Empty, error) {
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
	//mysql中的数据存储成功之后 , 再去存redis
	err := l.svcCtx.RDB.Pipelined(func(pipeliner redis.Pipeliner) error {
		for _, userId := range in.UserIds {
			if err := l.SaveGroupRedis(message, userId); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &chat.Empty{}, nil
}

// SaveGroupRedis redis 保存
func (l *SaveGroupMessageRedisLogic) SaveGroupRedis(message *types.GroupMessage, userId uint64) (err error) {
	score, _ := strconv.Atoi(message.Time)
	cache, _ := json.Marshal(message)
	_, err = l.svcCtx.RDB.Zadd(fmt.Sprintf("link:chat:user:%d", userId), int64(score), string(cache))
	return
}
