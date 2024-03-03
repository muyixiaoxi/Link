package logic

import (
	"fmt"
	"user/restful/internal/types"
	"user/service/user"
)

// GroupChat 群聊
func (l *ChatWSLogic) GroupChat(message types.Message) {
	//查询群聊内的人员列表
	userList, err := l.svcCtx.UserRpc.UserListByGroup(l.ctx, &user.SelectUserListByGroupRequest{GroupId: message.To})
	if err != nil {
		return
	}
	fmt.Println("消息", message)
	for _, groupUser := range userList.UserList {
		if groupUser.Id == message.From {
			//将当前用户信息拼接
			message.FromAvatar = groupUser.Avatar
			message.FromUsername = groupUser.Username
			break
		}
	}
	//判断群聊内用户谁在线
	for _, groupUser := range userList.UserList {
		c, has := Clients[groupUser.Id]
		if groupUser.Id != message.From && has {
			//消息不能发送给自己
			fmt.Println("$$$$$$$$$$$$$$$在线消息转发$$$$$$$$$$$$$$$$$$", message)
			c.Conn.WriteJSON(message)
		} else if !has {
			//如果是处于离线状态的用户 给此用户创建一个主题
			l.WriteByConn(message, groupUser.Id)
		}
	}
}
