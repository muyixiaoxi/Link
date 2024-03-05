package tagLogin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"tag/common/response"
	"tag/common/validate"
	"tag/restful/internal/logic/tagLogin"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"
)

func CreateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTagRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := tagLogin.NewCreateTagLogic(r.Context(), svcCtx)
		resp, err := l.CreateTag(&req)
		if err != nil {
			//fmt.Printf("Error type: %T, value: %v\n", err, err)
			formError, _ := status.FromError(err)
			if formError.Code() == codes.AlreadyExists {
				response.Response(w, nil, response.CodeTagIsExists)
				return
			}
			if formError.Code() == codes.NotFound {
				response.Response(w, nil, response.CodeSystemNotExist)
				return
			}
			if formError.Code() == 899 {
				response.Response(w, nil, response.CodeTagMoreMax)
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
