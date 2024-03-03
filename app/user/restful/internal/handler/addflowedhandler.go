package handler

import (
	"net/http"
	"strings"
	"user/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func addFlowedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAppleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddFlowedLogic(r.Context(), svcCtx)
		err := l.AddFlowed(&req)
		if err != nil {
			if strings.Contains(err.Error(), "repeat addition") {
				response.Response(w, nil, response.CodeRepeatAddition)
				return
			}
			if strings.Contains(err.Error(), "group chat overload") {
				response.Response(w, nil, response.CodeGroupChatOverload)
				return
			}
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
