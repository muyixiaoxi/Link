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
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/refreshToken/:refreshToken",
				Handler: refreshTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/app/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
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
					Method:  http.MethodGet,
					Path:    "/apply",
					Handler: GetApplyForHandler(serverCtx),
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
			}...,
		),
		rest.WithPrefix("/app/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/createGroup",
					Handler: userGroup.UserCreateGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/groupUserList",
					Handler: userGroup.GroupUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/kickOut",
					Handler: userGroup.KickOutUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/queryGroupList",
					Handler: userGroup.QueryGroupListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/quitGroup",
					Handler: userGroup.QuitGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/selectDeatilGroup",
					Handler: userGroup.SelectDetailGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/selectHomeGroup",
					Handler: userGroup.HomeGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/selectMyGroup",
					Handler: userGroup.SelectMyGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/selectStrangerGroup",
					Handler: userGroup.SelectStrangerGroupHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateGroupInfo",
					Handler: userGroup.UpdateGroupInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateGroupRemark",
					Handler: userGroup.UpdateGroupRemarkHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/app/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWT},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/cancelTag",
					Handler: userTag.CancelTagHandler(serverCtx),
				},
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
				{
					Method:  http.MethodPost,
					Path:    "/recommendUserList",
					Handler: userTag.RecommendUserListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/app/user"),
	)
}
