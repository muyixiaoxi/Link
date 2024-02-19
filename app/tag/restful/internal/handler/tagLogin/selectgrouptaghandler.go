package tagLogin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"tag/common/response"
	"tag/restful/internal/logic/tagLogin"
	"tag/restful/internal/svc"
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
