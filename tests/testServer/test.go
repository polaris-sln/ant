package main

import (
	"github.com/tor4z/ant/app"
	"github.com/tor4z/ant/web"
)

var routes web.Routes = web.Routes{
	"/hello": new(HelloHandler),
}

type HelloHandler struct{
	web.BaseHandler
}

func (h *HelloHandler)Get() {
	h.Send("hello")
}

func main() {
	app := app.NewApp()
	serve := web.NewServe(":8080")
	serve.SetRoute(routes)
	app.Run(serve)
}
