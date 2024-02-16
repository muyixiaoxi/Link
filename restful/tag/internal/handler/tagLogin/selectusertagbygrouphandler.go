package tagLogin

import (
	"Link/internal/response"
	"net/http"

	"Link/restful/tag/internal/logic/tagLogin"
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SelectUserTagByGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectUserTagByGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tagLogin.NewSelectUserTagByGroupLogic(r.Context(), svcCtx)
		resp, err := l.SelectUserTagByGroup(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)

		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
