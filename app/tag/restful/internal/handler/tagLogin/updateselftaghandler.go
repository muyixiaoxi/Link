package tagLogin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"tag/common/response"
	"tag/common/validate"
	"tag/restful/internal/logic/tagLogin"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSelfTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSelfTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			return
		}
		validatorErr := validate.Validate(&req)
		if validatorErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			logc.Error(context.Background(), "UpdateSelfTagHandler Params Error ", validatorErr)
			return
		}
		l := tagLogin.NewUpdateSelfTagLogic(r.Context(), svcCtx)
		err := l.UpdateSelfTag(&req)
		if err != nil {
			formError, _ := status.FromError(err)
			if formError.Code() == codes.NotFound {
				response.Response(w, nil, response.CodeTagNoExist)
				return
			}
			response.Response(w, nil, response.CodeServerBusy)
			logc.Error(context.Background(), "l.UpdateSelfTag(&req) is Failed", err)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
