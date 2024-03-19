package logic

import (
	"chat/restful/internal/types"
	"chat/service/user/user"
)

// GroupChat 群聊
func (l *ChatWSLogic) GroupChat(message types.Message) {
	//查询群聊内的人员列表
	userList, err := l.svcCtx.UserRpc.UserListByGroup(l.ctx, &user.SelectUserListByGroupRequest{GroupId: message.To})
	if err != nil {
		return
	}
	//判断群聊内用户谁在线
	for _, groupUser := range userList.UserList {
		c, has := Clients.Load(groupUser.Id)
		if groupUser.Id != message.From && has {
			//消息不能发送给自己
			err := c.(*Client).Conn.WriteJSON(message)
			if err != nil {
				//如果发生错误,也可能是中途离线了,此时也将消息存入kafka中
				WriteByConn(message, groupUser.Id)
			}
		} else if !has {
			//如果是处于离线状态的用户 给此用户创建一个主题
			WriteByConn(message, groupUser.Id)
		}
	}
}
