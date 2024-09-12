package manager

import (
	"money-manager/internal/constant"
	"money-manager/internal/schema"
	"net/http"
)

func GetRoutes() schema.RouteDetail {
	return schema.RouteDetail{
		Routing: schema.Routing{
			Version: constant.V1,
			Group:   constant.Manager,
		},
		Routes: []schema.Route{
			{
				ServiceName: constant.ManagerServiceAdd,
				Method:      http.MethodPost,
			},
		},
	}

}
