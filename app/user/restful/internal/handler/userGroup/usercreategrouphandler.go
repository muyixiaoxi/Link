package userGroup

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"net/http"
	"user/common/response"
	"user/common/validate"
	"user/restful/internal/logic/userGroup"
	"user/restful/internal/svc"
	"user/restful/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserCreateGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateGroupRequset
		httpx.Parse(r, &req)
		//参数校验
		validateErr := validate.Validate(req)
		if validateErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			logc.Error(context.Background(), "UserCreateGroupHandler Params is Error", validateErr)
			return
		}
		l := userGroup.NewUserCreateGroupLogic(r.Context(), svcCtx)
		err := l.UserCreateGroup(&req)
		if err != nil {
			fromErr, _ := status.FromError(err)
			if fromErr.Code() == 899 {
				//加入群聊达到最大限制
				response.Response(w, nil, response.CodeGroupMax)
				return
			}
			logc.Error(context.Background(), "userGroup.NewUserCreateGroupLogic(r.Context(), svcCtx) is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
