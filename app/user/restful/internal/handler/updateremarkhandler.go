package handler

import (
	"net/http"
	"user/common/response"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateRemarkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateRemarkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateRemarkLogic(r.Context(), svcCtx)
		err := l.UpdateRemark(&req)
		if err != nil {
			response.Response(w, nil, response.CodeFriendNotExist)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
