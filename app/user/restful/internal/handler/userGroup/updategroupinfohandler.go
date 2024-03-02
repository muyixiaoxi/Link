package userGroup

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"net/http"
	"user/common/response"
	"user/common/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic/userGroup"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func UpdateGroupInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGroupInfoRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			logc.Error(context.Background(), "UpdateGroupInfoHandler Params is failed", validateErr)
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := userGroup.NewUpdateGroupInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateGroupInfo(&req)
		if err != nil {
			fromErr, _ := status.FromError(err)
			if fromErr.Code() == 899 {
				response.Response(w, nil, response.CodeNoKick)
				return
			}
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
