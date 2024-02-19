package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"user/common/response"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			response.Response(w, resp, response.CodeUserNotExit)
			return
		}
		response.Response(w, resp, response.CodeSuccess)
	}
}
