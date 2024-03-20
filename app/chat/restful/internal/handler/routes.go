// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"chat/restful/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/chat",
				Handler: chatWSHandler(serverCtx),
			},
		},
		rest.WithPrefix("/app"),
	)
}
