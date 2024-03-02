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
	//判断群聊内用户谁在线
	for _, groupUser := range userList.UserList {
		c, has := Clients[groupUser.Id]
		if message.From == groupUser.Id {
			fmt.Println("消息不能发送给我自己:", groupUser.Id)
			message.FromAvatar = groupUser.Avatar
			message.FromUsername = groupUser.Username
		} else if groupUser.Id != message.From && has {
			//消息不能发送给自己
			c.Conn.WriteJSON(message)
		}
		//如果是处于离线状态的用户
	}
}
