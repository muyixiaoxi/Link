// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	userGroup "user/restful/internal/handler/userGroup"
	userTag "user/restful/internal/handler/userTag"
	"user/restful/internal/svc"

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
				Method:  http.MethodGet,
				Path:    "/chat",
				Handler: chatWSHandler(serverCtx),
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
				Method:  http.MethodPut,
				Path:    "/",
				Handler: updateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: getUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/flowed",
				Handler: addFlowedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/flowed/dispose",
				Handler: disposeFlowedHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/friend",
				Handler: DeleteFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friends",
				Handler: GetFriendsListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/friends/query",
				Handler: QueryFriendsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/query/phone",
				Handler: QueryPhoneHandler(serverCtx),
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
			{
				Method:  http.MethodPost,
				Path:    "/selectHomeGroup",
				Handler: userGroup.HomeGroupHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/app/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/chooseTag",
				Handler: userTag.UserChooseTagHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/queryLinkTags/:id",
				Handler: userTag.QueryLinkTagsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/app/user"),
	)
}
