package tagLogin

import (
	"Link/internal/response"
	"net/http"

	"Link/restful/tag/internal/logic/tagLogin"
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tagLogin.NewCreateTagLogic(r.Context(), svcCtx)
		resp, err := l.CreateTag(&req)
		if err != nil {
			if err.Error() == "TAG IS EXISTS" {
				response.Response(w, nil, response.CodeTagIsExists)
				return
			}
			response.Response(w, nil, response.CodeServerBusy)
			return
		}
		response.Response(w, resp, response.CodeSuccess)
		return
	}
}
