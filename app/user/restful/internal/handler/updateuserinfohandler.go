package handler

import (
	"net/http"
	"user/common/response"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateUserInfoLogic(r.Context(), svcCtx)
		err := l.UpdateUserInfo(&req)
		if err != nil {
			response.Response(w, nil, response.CodeUserExist)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
