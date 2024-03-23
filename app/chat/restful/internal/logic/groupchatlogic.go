package logic

import (
	"chat/restful/internal/types"
	"chat/service/chat"
	"chat/service/user/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GroupChat 群聊
func (l *ChatWSLogic) GroupChat(message types.Message) error {
	//1 . 查询全部的群成员
	groupUserList, err := l.svcCtx.UserRpc.UserListByGroup(l.ctx, &user.SelectUserListByGroupRequest{GroupId: message.To})
	if err != nil {
		return err
	}
	// 2 . 先向mysql内存入一份消息
	if err := l.SaveGroupMessage(message, nil); err != nil {
		return status.Error(codes.FailedPrecondition, "mysql存储失败")
	}
	// 3. 转发 , 然后获取到转发失败的用户的id
	var failedUserIds []uint64         //转发失败的群聊用户id
	onlineMap := map[string][]uint64{} //在其他服务端上在线的用户，分类
	var onlineUsers []uint64           // 在其他服务端上在线的所有用户
	for _, userInfo := range groupUserList.UserList {
		//不能转发给发送消息的用户
		userId := userInfo.Id
		// 群聊消息不能在转发一份给自己
		if userId != message.From {
			if c, has := Clients.Load(userId); has {
				if err := c.(*Client).Conn.WriteJSON(message); err != nil {
					//如果转发失败了 , 强制让当前用户断线
					l.Offline(userId)
					Clients.Delete(userId)
					//该用户属于转发失败的用户
					failedUserIds = append(failedUserIds, userId)
					return err
				}
			} else {
				//离线用户 首先判断是否在其他服务端在线
				resp, _ := l.svcCtx.ChatRpc.GetConnectorId(l.ctx, &chat.UserId{
					UserId: userId,
				})
				if resp != nil && resp.ConnectorId != "" {
					onlineMap[resp.ConnectorId] = append(onlineMap[resp.ConnectorId], userId)
					onlineUsers = append(onlineUsers, userId)
					continue
				}
				// 离线用户
				failedUserIds = append(failedUserIds, userId)
			}
		}
	}

	// 4. 转发其他在线用户的消息,转发失败当作离线处理
	err = Producer(onlineMap, message)
	if err != nil {
		failedUserIds = append(failedUserIds, onlineUsers...)
	}

	// 5 . 根据拿到的离线用户id 以及转发失败的用户id 将对应消息写入到他们的redis队列中
	if err := l.SaveGroupMessageRedis(message, failedUserIds); err != nil {
		return err
	}
	return nil
}
