package handler

import (
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
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		validatorErr := validate.Validate(&req)
		if validatorErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := logic.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
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
			response.Response(w, resp, response.CodeUserExist)
			return
		}
		response.Response(w, resp, response.CodeSuccess)
	}
}