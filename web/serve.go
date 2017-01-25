package web

import (
	"net/http"
	"sync"
	"time"
	"fmt"
	"github.com/tor4z/ant/app"
)

type Serve struct {
	addr     string
	keyFile  string
	certFile string
	tls      bool
	router   *Router
	app      *app.App
	mux      sync.Mutex
}

func NewServeTLS(addr string, certFile string, keyFile string) *Serve {
	return &Serve{addr:     addr,
		      certFile: certFile,
		      keyFile:  keyFile,
		      tls:      true,
		      router:   nil,
		      app:      nil,
	       }
}

func NewServe(addr string) *Serve{
	return &Serve{addr:     addr,
		      certFile: "",
		      keyFile:  "",
		      tls:      false,
		      router:   nil,
		      app:      nil,
	       }
}

func (serve *Serve)SetRoute(routes Routes) {
	serve.router = NewRouter(routes)
}

func (serve *Serve)SetApp(app *app.App) {
	serve.app = app
}

func (serve *Serve)Route(rw http.ResponseWriter, req *http.Request) {
	serve.mux.Lock()
	defer serve.mux.Unlock()
	handlerJob := NewHandlerJob(serve.router.Route(rw, req))
	handlerJob.Do()
	msg := fmt.Sprintf("[%s] %s %s", handlerJob.handler.Method(),
			   time.Now().Format("2006-01-02 15:04:05"),
			   req.URL.Path)
	go serve.app.LogInfo(msg)
}

func (serve *Serve)Do() {
	http.HandleFunc("/", serve.Route)
	fmt.Println("serve start ")
	if serve.tls {
		http.ListenAndServeTLS(serve.addr, serve.certFile,
		                       serve.keyFile, nil)
	} else {
		http.ListenAndServe(serve.addr, nil)
	}
	fmt.Println("serve exit")
}

