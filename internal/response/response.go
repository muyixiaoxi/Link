package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(w http.ResponseWriter, res any, err error) {
	if err != nil {
		body := Body{
			Code: -1,
			Data: nil,
			Msg:  err.Error(),
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}
	body := Body{
		Code: 0,
		Data: res,
		Msg:  "success",
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
