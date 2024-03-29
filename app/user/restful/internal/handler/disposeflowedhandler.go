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

func disposeFlowedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DisposeFlowedRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDisposeFlowedLogic(r.Context(), svcCtx)
		err := l.DisposeFlowed(&req)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				response.Response(w, nil, response.CodeApplyRecordNotExist)
			} else {
				response.Response(w, nil, response.CodeServerBusy)
			}
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
