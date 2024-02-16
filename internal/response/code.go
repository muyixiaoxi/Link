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
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:        "success",
	CodeParamsError:    "请求参数错误",
	CodeServerBusy:     "服务繁忙",
	CodeTagIsExists:    "标签已经存在",
	CodeUserExist:      "用户已存在",
	CodeUserNotExit:    "用户不存在",
	CodeSystemNotExist: "系统标签不存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
