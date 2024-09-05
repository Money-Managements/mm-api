package server

import (
	"money-manager/core/account"
	"money-manager/core/location"
	"money-manager/internal/schema"
	"net/http"

	"github.com/labstack/echo/v4"
)

func setRoutes(e *echo.Echo) {
	routes := getRoutes()
	groupApi, groupVersionMap, groupCoreMap := getGroupCollections(e)

	for _, route := range routes {
		groupVersion, _ := setGroupRoute(route.Version, groupApi, groupVersionMap)
		groupCore, _ := setGroupRoute(route.Group, groupVersion, groupCoreMap)
		setRouteMethod(route, groupCore)
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
		routeSettingBase := getRoutes()
		routeSetting := setRouteSetting(routeSettingBase.Routes, routeSettingBase.Routing)
		routes = append(routes, routeSetting...)
	}

	return routes
}

func setRouteSetting(routes []schema.Route, routing schema.Routing) []schema.Route {
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

func setRouteMethod(route schema.Route, groupCore *echo.Group) {
	switch route.Method {
	case http.MethodGet:
		groupCore.GET(route.EndPoint, route.Handler)
	case http.MethodPost:
		groupCore.POST(route.EndPoint, route.Handler)
	case http.MethodPatch:
		groupCore.PUT(route.EndPoint, route.Handler)
	case http.MethodDelete:
		groupCore.DELETE(route.EndPoint, route.Handler)
	default:
		println("Unsupported method: " + route.Method)
	}
}

func logSettedRoute(route schema.Route) {
	println(route.Method + " : " + route.Version + "/" + route.Group + route.EndPoint)
}
