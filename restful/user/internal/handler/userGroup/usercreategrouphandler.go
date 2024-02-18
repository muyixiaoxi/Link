package userGroup

import (
	"Link/internal/response"
	"Link/internal/validate"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"

	"Link/restful/user/internal/logic/userGroup"
	"Link/restful/user/internal/svc"
	"Link/restful/user/internal/types"
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
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
