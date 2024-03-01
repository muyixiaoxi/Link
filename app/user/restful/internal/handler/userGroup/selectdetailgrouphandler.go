package userGroup

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"user/common/response"
	"user/common/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic/userGroup"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func SelectDetailGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectGroupDeatilRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			logc.Error(context.Background(), "SelectDetailGroupHandler Params is error", validateErr)
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := userGroup.NewSelectDetailGroupLogic(r.Context(), svcCtx)
		resp, err := l.SelectDetailGroup(&req)
		if err != nil {
			logc.Error(context.Background(), "userGroup.NewSelectDetailGroupLogic(r.Context(), svcCtx) is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
