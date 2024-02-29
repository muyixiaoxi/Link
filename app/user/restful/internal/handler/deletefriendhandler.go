package handler

import (
	"net/http"
	"user/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func DeleteFriendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDeleteFriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteFriendLogic(r.Context(), svcCtx)
		err := l.DeleteFriend(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
