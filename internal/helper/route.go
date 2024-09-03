package helper

import "money-manager/internal/schema"

func SetRoute(routeDetail schema.RouteDetail) []schema.Route {
	return SetRouteSetting(routeDetail.Routes, routeDetail.Routing)
}

func SetRouteSetting(routes []schema.Route, routing schema.Routing) []schema.Route {
	for i := range routes {
		if routes[i].Group == "" {
			routes[i].Group = routing.Group
		}

		if routes[i].Version == "" {
			routes[i].Version = routing.Version
		}

		if routes[i].EndPoint == "/" {
			routes[i].EndPoint = ""
		}
	}

	return routes
}
