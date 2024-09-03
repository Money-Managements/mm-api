package schema

import "github.com/labstack/echo/v4"

type Route struct {
	Routing
	Method   string
	EndPoint string
	Handler  echo.HandlerFunc
}

type Routing struct {
	Version string
	Group   string
}

type RouteDetail struct {
	Routing
	Routes []Route
}
