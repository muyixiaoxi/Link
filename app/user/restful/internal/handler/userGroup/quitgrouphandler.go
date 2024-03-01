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

func QuitGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuitGroupRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			logc.Error(context.Background(), "QuitGroupHandler Params is error", validateErr)
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := userGroup.NewQuitGroupLogic(r.Context(), svcCtx)
		err := l.QuitGroup(&req)
		if err != nil {
			logc.Error(context.Background(), "userGroup.NewQuitGroupLogic(r.Context(), svcCtx) is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
