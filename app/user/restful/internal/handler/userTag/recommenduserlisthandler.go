package userTag

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"
	"user/common/response"
	"user/common/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/restful/internal/logic/userTag"
	"user/restful/internal/svc"
	"user/restful/internal/types"
)

func RecommendUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendGroupByTagRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			logc.Error(context.Background(), "RecommendUserListHandler params is error", validateErr)
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := userTag.NewRecommendUserListLogic(r.Context(), svcCtx)
		resp, err := l.RecommendUserList(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
