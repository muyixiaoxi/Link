package response

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeParamsError
	CodeServerBusy
	CodeTagIsExists
	CodeUserExist
	CodeUserNotExit
	CodeSystemNotExist
	CodeTagNoExist
	CodeTagSame
	CodeTagChooseFailed
	CodeWebSocketLinkFail
	CodeFriendNotExist
	CodeTagMoreMax
	CodeTokenFormat
	CodeTokenIsEmpty
	CodeTokenInvalid
	CodeTokenPast
	CodeTokenIsFailed
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeParamsError:       "请求参数错误",
	CodeServerBusy:        "服务繁忙",
	CodeTagIsExists:       "标签已经存在",
	CodeUserExist:         "用户已存在",
	CodeUserNotExit:       "用户不存在",
	CodeSystemNotExist:    "系统标签不存在",
	CodeTagNoExist:        "标签不存在",
	CodeTagSame:           "标签被重复选择",
	CodeTagChooseFailed:   "标签选择失败",
	CodeWebSocketLinkFail: "WebSocket连接失败",
	CodeFriendNotExist:    "好友不存在",
	CodeTagMoreMax:        "标签超过最大数量限制",
	CodeTokenFormat:       "令牌格式错误",
	CodeTokenIsEmpty:      "令牌为空",
	CodeTokenInvalid:      "令牌无效,请重新登录",
	CodeTokenPast:         "登录过期",
	CodeTokenIsFailed:     "Token签发失败",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
