package handler

import (
	"net/http"

	"Link/restful/tag/internal/logic"
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTagLogic(r.Context(), svcCtx)
		resp, err := l.Tag(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
