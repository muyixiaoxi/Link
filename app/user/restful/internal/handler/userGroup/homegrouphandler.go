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

func HomeGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendGroupByTagRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			logc.Error(context.Background(), "HomeGroupHandler params is error", validateErr)
			return
		}
		l := userGroup.NewHomeGroupLogic(r.Context(), svcCtx)
		resp, err := l.HomeGroup(&req)
		if err != nil {
			logc.Error(context.Background(), " userGroup.NewHomeGroupLogic(r.Context(), svcCtx) is failed", err)
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
