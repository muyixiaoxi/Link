package handler

import (
	"net/http"
	"user/common/response"

	"user/restful/internal/logic"
	"user/restful/internal/svc"
)

func GetFriendsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetFriendsListLogic(r.Context(), svcCtx)
		resp, err := l.GetFriendsList()
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
