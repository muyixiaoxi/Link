package tagLogin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tag/common/response"
	"tag/common/validate"
	"tag/restful/internal/logic/tagLogin"
	"tag/restful/internal/svc"
	"tag/restful/internal/types"
)

func SelectUserTagByGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectUserTagByGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			return
		}
		validateErr := validate.Validate(req)
		if validateErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			logc.Error(context.Background(), "SelectUserTagByGroupHandler validateErr := validate.Validate(req) is failed ", validateErr)
			return
		}
		l := tagLogin.NewSelectUserTagByGroupLogic(r.Context(), svcCtx)
		resp, err := l.SelectUserTagByGroup(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
			logc.Error(context.Background(), "SelectUserTagByGroupHandler Params Is Error ", err)
		} else {
			response.Response(w, resp, response.CodeSuccess)
		}
	}
}
