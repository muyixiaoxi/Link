// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	userGroup "Link/restful/user/internal/handler/userGroup"
	"Link/restful/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: signUpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/app/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: getUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/",
				Handler: updateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/flowed",
				Handler: addFlowedHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/remark",
				Handler: updateRemarkHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/app/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/createGroup",
				Handler: userGroup.UserCreateGroupHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/app/user"),
	)
}
