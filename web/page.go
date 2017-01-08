package web

import (
	"net/http"
)

type Page struct {
	rw  http.ResponseWriter
	req http.Request
}

func (page *Page)send() {


}
