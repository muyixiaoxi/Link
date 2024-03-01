package handler

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"user/common/response"
	"user/common/validate"
	"user/restful/internal/logic"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func signUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateRequest
		httpx.Parse(r, &req)
		validatorErr := validate.Validate(&req)
		if validatorErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := logic.NewSignUpLogic(r.Context(), svcCtx)
		_, err := l.SignUp(&req)
		fmt.Println("@@", err)
		if err != nil {
			formError, _ := status.FromError(err)
			if formError.Code() == codes.AlreadyExists {
				response.Response(w, nil, response.CodeUserExist)
				return
			}
			logc.Error(context.Background(), "signUpHandler is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
			return
		}
		response.Response(w, nil, response.CodeSuccess)
	}
}
