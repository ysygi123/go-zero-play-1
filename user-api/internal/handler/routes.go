// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-play-1/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.TestMiddleware1},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/user/get/:id",
					Handler: getUserHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/user/jwt",
				Handler: getJWTHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/yys/setYys",
				Handler: setYysHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/yys/calculateInventory",
				Handler: calculateInventoryHandler(serverCtx),
			},
		},
	)
}
