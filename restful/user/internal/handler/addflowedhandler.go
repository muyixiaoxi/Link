package handler

import (
	"Link/internal/response"
	"net/http"

	"Link/restful/user/internal/logic"
	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func addFlowedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFlowedRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddFlowedLogic(r.Context(), svcCtx)
		err := l.AddFlowed(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
