package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {

	router := NewRouter()

	router.Handle("GET", "/", IndexHandler)
	router.Handle("GET", "/hello", HelloHandler)
	router.Handle("POST", "/echo", EchoHandler)

	staticHandler := fasthttp.FSHandler("./static", 0)
	router.Handle("GET", "/static/*filepath", staticHandler)

	// Start the server
	if err := fasthttp.ListenAndServe(":8080", router.HandleRequest); err != nil {
		fmt.Printf("Error in ListenAndServe: %s\n", err)
	}
}

type Router struct {
	routes map[string]map[string]fasthttp.RequestHandler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]fasthttp.RequestHandler),
	}
}

func (r *Router) Handle(method, path string, handler fasthttp.RequestHandler) {
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = make(map[string]fasthttp.RequestHandler)
	}
	r.routes[method][path] = handler
}

func (r *Router) HandleRequest(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	method := string(ctx.Method())

	if handler, ok := r.routes[method][path]; ok {
		handler(ctx)
	} else {
		ctx.Error("Not Found", fasthttp.StatusNotFound)
	}
}

//Handling requests to different paths

func IndexHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Index Page :o")
}

func HelloHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Hello :3")
}

func EchoHandler(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	ctx.Write(body)
}
