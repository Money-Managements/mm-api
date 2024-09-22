package schema

import "github.com/labstack/echo/v4"

type Route struct {
	Routing
	Handler  func(c echo.Context) error
	Method   string
	EndPoint string
}

type Routing struct {
	Version string
	Group   string
}

type RouteDetail struct {
	Routing
	Routes []Route
}
