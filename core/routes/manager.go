package routes

import (
	"money-manager/internal/constant"
	"money-manager/internal/schema"
	"net/http"
)

func GetManagerRoutes() schema.RouteDetail {
	return schema.RouteDetail{
		Routing: schema.Routing{
			Version: constant.V1,
			Group:   constant.Manager,
		},
		Routes: []schema.Route{
			{
				Method: http.MethodPost,
			},
		},
	}

}
