package server

import (
	"money-manager/core/account"
	"money-manager/internal/schema"

	"github.com/labstack/echo/v4"
)

func setRoutes(e *echo.Echo) {
	var routes []schema.Route
	routes = append(routes, account.GetRoutes()...)

	groupApi := e.Group("/api")
	groupVersionMap := make(map[string]*echo.Group)
	groupCoreMap := make(map[string]*echo.Group)

	for _, route := range routes {
		// Ensure the groupVersion is created and assigned properly
		groupVersion, existsGroupVersion := groupVersionMap[route.Version]
		if !existsGroupVersion {
			groupVersion = groupApi.Group("/" + route.Version)
			groupVersionMap[route.Version] = groupVersion
		}

		// Ensure the groupCore is created and assigned properly
		groupCore, existsGroupCore := groupCoreMap[route.Group]
		if !existsGroupCore {
			groupCore = groupVersion.Group("/" + route.Group)
			groupCoreMap[route.Group] = groupCore
		}

		groupCore.GET(route.EndPoint, route.Handler)
		println(route.Method + " : " + route.Group + "/" + route.Version + route.EndPoint)
	}
}
