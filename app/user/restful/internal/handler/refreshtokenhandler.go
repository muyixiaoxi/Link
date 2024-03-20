package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"user/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func refreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshTokenRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken(&req)
		if err != nil {
			fromErr, _ := status.FromError(err)
			if fromErr.Code() == codes.NotFound || fromErr.Code() == codes.FailedPrecondition {
				response.Response(w, nil, response.CodeTokenIsEmpty)
				return
			}
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
