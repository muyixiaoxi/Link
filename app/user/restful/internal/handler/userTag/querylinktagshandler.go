package userTag

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"user/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic/userTag"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func QueryLinkTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := userTag.NewQueryLinkTagsLogic(r.Context(), svcCtx)
		resp, err := l.QueryLinkTags(&req)
		if err != nil {
			logc.Error(context.Background(), "l.QueryLinkTags() is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
