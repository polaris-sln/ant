package web

import (
	"net/http"
	"io"
	"regexp"
)

type HandlerJob struct {
	handler Handler
}

type Handler interface {
	Get()
	Post()
	Put()
	Head()
	Trace()
	Delete()
	Option()
	Connect()
	Patch()
	Method() string
	MethodNotAllow()
	SetResponseWriter(rw http.ResponseWriter)
	SetRequest(req *http.Request)
	Send(msg string)
	SetStatusCode(code int)
}

func NewHandlerJob(handler Handler) *HandlerJob {
	return &HandlerJob{handler,}
}

func (hr *HandlerJob)Do() {
	method := hr.handler.Method()
	switch method {
	case "Get":     hr.handler.Get()
	case "POST":    hr.handler.Post()
	case "PUT":     hr.handler.Put()
	case "HEAD":    hr.handler.Head()
	case "DELETE":  hr.handler.Delete()
	case "TRACE":   hr.handler.Trace()
	case "OPTION":  hr.handler.Option()
	case "CONNECT": hr.handler.Connect()
	case "PATCH":   hr.handler.Patch()
	default:        hr.handler.MethodNotAllow()
	}
}


type BaseHandler struct {
	responseWriter http.ResponseWriter
	request        *http.Request
}

func (handler *BaseHandler)SetResponseWriter(rw http.ResponseWriter) {
	handler.responseWriter = rw
}

func (handler *BaseHandler)SetRequest(req *http.Request) {
	handler.request = req
}

func (handler *BaseHandler)Get() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Post() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Head() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Trace() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Put() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Delete() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Option() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Connect() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)Patch() {
	handler.MethodNotAllow()
}

func (handler *BaseHandler)MethodNotAllow() {
	html := `<html>
		<head><title>405 Not Allowed</title></head>
		<body bgcolor='white'>
		<center><h1>405 Not Allowed</h1></center>
		<hr><center>ant</center>
		</body>
		</html>`
	handler.SetStatusCode(http.StatusMethodNotAllowed)
	handler.Send(html)
}

func (handler *BaseHandler)Method() string {
	return handler.request.Method
}

func (handler *BaseHandler)Redirect(path string) {
	http.Redirect(handler.responseWriter,
	              handler.request, path,
		      http.StatusFound)
}

func (handler *BaseHandler)Send(str string) {
	io.WriteString(handler.responseWriter, str)
}

func (handler *BaseHandler)SetStatusCode(code int) {
	handler.responseWriter.WriteHeader(code)
}

func (handler *BaseHandler)SetHeader(key string, value string){
	handler.responseWriter.Header().Set(key, value)
}

type Router struct {
	routes Routes
}

func NewRouter(routes Routes) *Router {
	return &Router{routes,}
}

func (router *Router)SetRoutes(routes Routes) {
	router.routes = routes
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


type Server struct {
	addr    string
	keyFile string
	certFile string
	tls     bool
}

func NewServer(addr string, tls bool, keyFile ...string) *Server{
	var kf string
	var cf string

	if len(keyFile) == 1{
		kf = keyFile[0]
		cf = keyFile[1]
	} else {
		kf = ""
		cf = ""
	}
	return &Server{addr, kf, cf, tls,}
}

func (server *Server)Do() {
	if server.tls {
		http.ListenAndServeTLS(server.addr, server.certFile, server.keyFile, nil)
	} else {
		http.ListenAndServe(server.addr, nil)
	}
}

