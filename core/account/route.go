package account

import (
	"money-manager/internal/constant"
	"money-manager/internal/schema"
	"net/http"
)

func GetRoutes() schema.RouteDetail {
	return schema.RouteDetail{
		Routing: schema.Routing{
			Version: constant.V1,
			Group:   constant.Account,
		},
		Routes: []schema.Route{
			{
				ServiceName: constant.AcccountServiceGet,
				Method:      http.MethodGet,
			},
		},
	}

}
