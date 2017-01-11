package web

import (
	"net/http"
	"regexp"
)

type Routes map[string]Handler

type Router struct {
	routes Routes
}

func NewRouter(routes Routes) *Router {
	return &Router{routes:routes,}
}

func (router *Router)Route(rw http.ResponseWriter, req *http.Request) Handler {
	for pattern, handler := range router.routes {
		re := regexp.MustCompile(pattern)
		if re.MatchString(req.URL.Path) {
			handler.SetResponseWriter(rw)
			handler.SetRequest(req)
			return handler
		}
	}
	handler := new(NotFoundHandler)
	handler.SetResponseWriter(rw)
	handler.SetRequest(req)
	return handler
}


