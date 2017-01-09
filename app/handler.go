package app

import (
	"net/http"
	"io"
)

type Handler struct {
	responseWriter http.ResponseWriter
	request        http.Request
}

func (handler *Handler)Do() {

}

func (handler *Handler)Get() {
	handler.MethodNotAllow()
}

func (handler *Handler)Post() {
	handler.MethodNotAllow()
}

func (handler *Handler)Head() {
	handler.MethodNotAllow()
}

func (handler *Handler)Trace() {
	handler.MethodNotAllow()
}

func (handler *Handler)Put() {
	handler.MethodNotAllow()
}

func (handler *Handler)Delete() {
	handler.MethodNotAllow()
}

func (handler *Handler)Option() {
	handler.MethodNotAllow()
}

func (handler *Handler)Connect() {
	handler.MethodNotAllow()
}

func (handler *Handler)Patch() {
	handler.MethodNotAllow()
}

func (handler *Handler)MethodNotAllow() {
	html := "<html>
		<head><title>405 Not Allowed</title></head>
		<body bgcolor='white'>
		<center><h1>405 Not Allowed</h1></center>
		<hr><center>ant</center>
		</body>
		</html>"
	handler.SetStatusCode(http.StatusMethodNotAllowed)
	handler.Send(html)
}

func (handler *Handler)Send(str string) {
	io.WriteString(handler.responseWriter, str)
}

func (handler *Handler)SetStatusCode(code int) {
	handler.responseWriter.WriteHeader(code)
}

func (handler *Handler)SetHeadler(key string, value string){
	handler.responseWriter.Header().Set(key, value)
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

