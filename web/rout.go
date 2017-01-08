package web

import (
	"net/http"
)

type Route map[string]http.HandlerFunc
