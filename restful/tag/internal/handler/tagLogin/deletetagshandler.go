package tagLogin

import (
	"Link/internal/response"
	"Link/internal/validate"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"

	"Link/restful/tag/internal/logic/tagLogin"
	"Link/restful/tag/internal/svc"
	"Link/restful/tag/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelteTags
		if err := httpx.Parse(r, &req); err != nil {
			return
		}
		//参数校验
		validatorErr := validate.Validate(&req)
		if validatorErr != nil {
			response.Response(w, nil, response.CodeParamsError)
			logc.Error(context.Background(), "DeleteTagsHandler Params is Error", validatorErr)
			return
		}

		l := tagLogin.NewDeleteTagsLogic(r.Context(), svcCtx)
		err := l.DeleteTags(&req)
		if err != nil {
			response.Response(w, nil, response.CodeServerBusy)
			logc.Error(context.Background(), "l.DeleteTags(&req) is Error", err)
		} else {
			response.Response(w, nil, response.CodeSuccess)
		}
	}
}
