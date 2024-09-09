package server

import (
	"money-manager/core/account"
	"money-manager/core/location"
	"money-manager/internal/schema"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setRoutes(e *echo.Echo, db *gorm.DB) {
	routes := getRoutes()
	groupApi, groupVersionMap, groupCoreMap := getGroupCollections(e)
	services := getServices(db)

	for _, route := range routes {
		groupVersion, _ := setGroupRoute(route.Version, groupApi, groupVersionMap)
		groupCore, _ := setGroupRoute(route.Group, groupVersion, groupCoreMap)
		route.Service = services[route.ServiceName].Action
		setRouteMethod(route, groupCore, route.Service)
		logSettedRoute(route)
	}
}

func getRoutes() []schema.Route {
	getRoutesFuncs := []func() schema.RouteDetail{
		account.GetRoutes,
		location.GetRoutes,
	}

	var routes []schema.Route

	for _, getRoutes := range getRoutesFuncs {
		routeDetail := getRoutes()
		routeSetting := setRouteSetting(routeDetail)
		routes = append(routes, routeSetting...)
	}

	return routes
}

func setRouteSetting(routeDetail schema.RouteDetail) []schema.Route {
	routes := routeDetail.Routes
	routing := routeDetail.Routing

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

func getGroupCollections(e *echo.Echo) (*echo.Group, map[string]*echo.Group, map[string]*echo.Group) {
	groupApi := e.Group("/api")
	groupVersionMap := make(map[string]*echo.Group)
	groupCoreMap := make(map[string]*echo.Group)
	return groupApi, groupVersionMap, groupCoreMap
}

func setGroupRoute(pathRoute string, groupBase *echo.Group, groupCollection map[string]*echo.Group) (*echo.Group, bool) {
	group, existGroup := groupCollection[pathRoute]
	if !existGroup {
		group = groupBase.Group("/" + pathRoute)
		groupCollection[pathRoute] = group
	}

	return group, existGroup
}

func setRouteMethod(route schema.Route, groupCore *echo.Group, service schema.ActionService) {
	handler := getHandler(service)

	switch route.Method {
	case http.MethodGet:
		groupCore.GET(route.EndPoint, handler)
	case http.MethodPost:
		groupCore.POST(route.EndPoint, handler)
	case http.MethodPatch:
		groupCore.PATCH(route.EndPoint, handler)
	case http.MethodDelete:
		groupCore.DELETE(route.EndPoint, handler)
	default:
		println("Unsupported method: " + route.Method)
	}
}

func getHandler(service schema.ActionService) func(c echo.Context) error {
	handler := func(c echo.Context) error {
		service(c)
		return c.String(http.StatusOK, "money-manager")
	}

	return handler
}

func logSettedRoute(route schema.Route) {
	println(route.Method + " : " + route.Version + "/" + route.Group + route.EndPoint)
}
