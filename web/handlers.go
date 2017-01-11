package web

import (
	"net/http"
)

type NotFoundHandler struct {
	BaseHandler
}

func (h *NotFoundHandler)Get() {
	html := `<html>
		<head><title>404 Not Found</title></head>
		<body bgcolor='white'>
		<center><h1>404 Not Found</h1></center>
		<hr><center>ant</center>
		</body>
		</html>`
	h.SetStatusCode(http.StatusNotFound)
	h.Send(html)
}

