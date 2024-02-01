package response

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeServerBusy
	CodeTagIsExists
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:     "success",
	CodeServerBusy:  "服务繁忙",
	CodeTagIsExists: "标签已经存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
