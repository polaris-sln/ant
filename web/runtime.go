package web

var Ant = new(ant)

func AntSetAddr(addr string) {
	Ant.addr = addr
}

func AntSetRoutes(routes Route) {
	Ant.routes = routes;
}
