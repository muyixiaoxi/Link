package tagLogin

import (
	"net/http"
	"tag/common/response"

	"tag/restful/internal/logic/tagLogin"
	"tag/restful/internal/svc"
)

func SelectAllTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tagLogin.NewSelectAllTagsLogic(r.Context(), svcCtx)
		resp, err := l.SelectAllTags()
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
