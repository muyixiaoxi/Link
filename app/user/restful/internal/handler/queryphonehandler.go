package handler

import (
	"net/http"
	"user/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func QueryPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserQueryPhoneRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewQueryPhoneLogic(r.Context(), svcCtx)
		resp, err := l.QueryPhone(&req)
		if err != nil {
			response.Response(w, nil, response.CodeUserNotExit)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
