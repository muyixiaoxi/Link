package tagLogin

import (
	"Link/internal/response"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"

	"Link/restful/tag/internal/logic/tagLogin"
	"Link/restful/tag/internal/svc"
)

func SelectGroupTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tagLogin.NewSelectGroupTagLogic(r.Context(), svcCtx)
		resp, err := l.SelectGroupTag()
		if err != nil {
			logc.Error(context.Background(), "l.SelectGroupTag() is error", err)
			response.Response(w, nil, response.CodeServerBusy)
			return
		}
		response.Response(w, resp, response.CodeSuccess)
	}
}
