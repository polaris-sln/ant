package app

import (
	"net/http"
	"regexp"
)

func antInit(addr string, routes Route) {
	AntSetAddr(addr)
	AntSetRoutes(routes)
}

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	for pattern, handler := range Ant.routes {
		re := regexp.MustCompile(pattern)
		if re.MatchString(req.URL.Path) {
			go handler(rw, req)
			return
		}
	}
	http.NotFound(rw, req)
}

func Start(addr string, routes Route) {
	antInit(addr, routes)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(addr, nil)
}
