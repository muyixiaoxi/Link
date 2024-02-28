package handler

import (
	"fmt"
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
			formErr, _ := status.FromError(err)
			switch formErr.Code() {
			case codes.AlreadyExists:
				//标签被重复选择
				response.Response(w, nil, response.CodeTagSame)
				return
			case codes.DeadlineExceeded:
				//标签选择失败
				response.Response(w, nil, response.CodeTagChooseFailed)
				return

			}
			response.Response(w, nil, response.CodeUserExist)
			return
		}
		response.Response(w, nil, response.CodeSuccess)
	}
}
