package main

import (
	"github.com/iaoizo/zinx/examples/zinx_interceptor/interceptors"
	"github.com/iaoizo/zinx/examples/zinx_interceptor/router"
	"github.com/iaoizo/zinx/znet"
)

func main() {
	server := znet.NewServer()

	server.AddRouter(1, &router.HelloRouter{})

	// Add Custom Interceptor
	server.AddInterceptor(&interceptors.MyInterceptor{})

	server.Serve()
}
