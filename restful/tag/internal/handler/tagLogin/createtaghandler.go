package tagLogin

import (
	"Link/internal/response"
	"Link/restful/tag/internal/logic/tagLogin"
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
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
			//fmt.Printf("Error type: %T, value: %v\n", err, err)
			formError, _ := status.FromError(err)
			if formError.Code() == codes.AlreadyExists {
				logc.Error(context.Background(), "TAG IS EXISTS")
				response.Response(w, nil, response.CodeTagIsExists)
				return
			}
			logc.Error(context.Background(), "tagLogin.NewCreateTagLogic(r.Context(), svcCtx) is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
			return
		}
		response.Response(w, resp, response.CodeSuccess)
		return
	}
}
