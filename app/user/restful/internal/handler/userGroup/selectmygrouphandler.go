package userGroup

import (
	"net/http"
	"user/common/response"
	"user/common/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic/userGroup"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func SelectMyGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectMyGroupRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(&req)
		if validateErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := userGroup.NewSelectMyGroupLogic(r.Context(), svcCtx)
		resp, err := l.SelectMyGroup(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
