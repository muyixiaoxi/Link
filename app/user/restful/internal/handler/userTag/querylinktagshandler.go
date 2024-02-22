package userTag

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"user/common/response"

	"user/restful/internal/logic/userTag"
	"user/restful/internal/svc"
)

func QueryLinkTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := userTag.NewQueryLinkTagsLogic(r.Context(), svcCtx)
		resp, err := l.QueryLinkTags()
		if err != nil {
			logc.Error(context.Background(), "l.QueryLinkTags() is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
