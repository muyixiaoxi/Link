package handler

import (
	"Link/internal/response"
	"Link/restful/user/internal/logic"
	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func signUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
		if err != nil {
			response.Response(w, resp, response.CodeUserNotExit)
		}
		response.Response(w, resp, response.CodeSuccess)
	}
}
