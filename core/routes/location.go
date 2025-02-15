package routes

import (
	"money-manager/internal/constant"
	"money-manager/internal/schema"
	"net/http"
)

func GetLocationRoutes() schema.RouteDetail {
	return schema.RouteDetail{
		Routing: schema.Routing{
			Version: constant.V1,
			Group:   constant.Location,
		},
		Routes: []schema.Route{
			{
				Method: http.MethodGet,
			},
		},
	}

}
