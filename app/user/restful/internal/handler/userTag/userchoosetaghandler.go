package userTag

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"user/common/response"
	"user/common/validate"
	"user/restful/internal/logic/userTag"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func UserChooseTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserChooseTagRequst
		httpx.Parse(r, &req)
		//验证参数
		validatorErr := validate.Validate(&req)
		if validatorErr != nil {
			logc.Error(context.Background(), "UserChooseTagHandler params is error", validatorErr)
			response.Response(w, nil, response.CodeParamsError)
			return
		}

		l := userTag.NewUserChooseTagLogic(r.Context(), svcCtx)
		err := l.UserChooseTag(&req)
		if err != nil {
			fromErr, _ := status.FromError(err)
			if fromErr.Code() == 899 {
				response.Response(w, nil, response.CodeTagMoreMax)
				return
			}
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
