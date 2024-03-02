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

func UpdateGroupRemarkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGroupRemarkRequest
		httpx.Parse(r, &req)
		validateErr := validate.Validate(req)
		if validateErr != nil {
			logc.Error(context.Background(), "UpdateGroupRemarkHandler Params is error", validateErr)
			response.Response(w, nil, response.CodeParamsError)
			return
		}
		l := userGroup.NewUpdateGroupRemarkLogic(r.Context(), svcCtx)
		err := l.UpdateGroupRemark(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
