package schema

type Route struct {
	Routing
	ServiceName string
	Service     ActionService
	Method      string
	EndPoint    string
}

type Routing struct {
	Version string
	Group   string
}

type RouteDetail struct {
	Routing
	Routes []Route
}
