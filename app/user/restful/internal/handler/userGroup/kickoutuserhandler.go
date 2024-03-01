package userGroup

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"user/common/response"
	"user/common/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic/userGroup"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func KickOutUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KickOutGroupRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			logc.Error(context.Background(), "KickOutUserHandler Params is error", validateErr)
			return
		}
		l := userGroup.NewKickOutUserLogic(r.Context(), svcCtx)
		err := l.KickOutUser(&req)
		if err != nil {
			formError, _ := status.FromError(err)
			if formError.Code() == codes.Unavailable {
				response.Response(w, nil, response.CodeNoKick)
				return
			}
			logc.Error(context.Background(), "userGroup.NewKickOutUserLogic(r.Context(), svcCtx) is error", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
