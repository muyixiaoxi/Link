package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code ResCode `json:"code"`
	Data any     `json:"data"`
	Msg  string  `json:"msg"`
}

func InitBody(res any, code ResCode) (body Body) {
	if code > 1000 {
		body = Body{
			Code: code,
			Data: nil,
			Msg:  code.Msg(),
		}
		return
	}
	body = Body{
		Code: code,
		Data: res,
		Msg:  code.Msg(),
	}
	return
}

func Response(w http.ResponseWriter, res any, code ResCode) {
	if code > 1000 {
		body := Body{
			Code: code,
			Data: nil,
			Msg:  code.Msg(),
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}
	body := Body{
		Code: code,
		Data: res,
		Msg:  code.Msg(),
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
