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

func SelectStrangerGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectStrangerRequest
		httpx.Parse(r, &req)
		if validateErr := validate.Validate(req); validateErr != nil {
			response.Response(w, validateErr, response.CodeParamsError)
			return
		}
		l := userGroup.NewSelectStrangerGroupLogic(r.Context(), svcCtx)
		resp, err := l.SelectStrangerGroup(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
