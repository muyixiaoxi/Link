package tagLogin

import (
	"net/http"

	"Link/restful/tag/internal/logic/tagLogin"
	"Link/restful/tag/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SelectGourpTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tagLogin.NewSelectGourpTagLogic(r.Context(), svcCtx)
		err := l.SelectGourpTag()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
