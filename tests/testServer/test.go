package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/tor4z/ant/web"
)

var routes web.Route = web.Route{
	"/hello": helloHandler,
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("helloS");
	io.WriteString(rw, "hello")
	fmt.Println("helloE");
}

func main() {
	fmt.Println("Testing");
	web.Start(":8080", routes)
	fmt.Println("end");
}
