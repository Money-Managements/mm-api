package routes

import (
	"money-manager/core/handlers"
	"money-manager/internal/constant"
	"money-manager/internal/schema"
	"net/http"
)

func GetAccountRoutes() schema.RouteDetail {
	return schema.RouteDetail{
		Routing: schema.Routing{
			Version: constant.V1,
			Group:   constant.Account,
		},
		Routes: []schema.Route{
			{
				Handler: handlers.GetAccountHandler,
				Method:  http.MethodGet,
			},
		},
	}

}
