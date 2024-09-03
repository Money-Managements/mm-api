package account

import (
	"money-manager/internal/constant"
	"money-manager/internal/helper"
	"money-manager/internal/schema"
	"net/http"
)

func GetRoutes() []schema.Route {
	routeSetting := schema.RouteDetail{
		Routing: schema.Routing{
			Version: constant.V1,
			Group:   constant.Account,
		},
		Routes: []schema.Route{
			{
				Method:  http.MethodGet,
				Handler: GetAccounts,
			},
		},
	}

	return helper.SetRoute(routeSetting)
}
