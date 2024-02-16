package tagTest

import (
	"net/http"

	"Link/restful/tag/internal/logic/tagTest"
	"Link/restful/tag/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tagTest.NewUserTagLogic(r.Context(), svcCtx)
		err := l.UserTag()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
