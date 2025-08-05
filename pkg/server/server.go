package server

import (
	"fmt"
	"net/http"

	"github.com/fernandojosemoran/go-templates/pkg/enums"
	"github.com/fernandojosemoran/go-templates/pkg/logger"
)

type Controller struct {
	Method  enums.Method
	Path    string
	Handler http.HandlerFunc
}

type HttpServer struct {
	Port        int
	Middlewares []http.HandlerFunc
	Handlers    []Controller
}

func startHandlers(c []Controller) {
	for _, handler := range c {
		http.HandleFunc(fmt.Sprintf("%s %s", handler.Method, handler.Path), handler.Handler.ServeHTTP)
	}
}

func middlewares(next http.Handler, middlewares []http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			middleware(w, r)
		}
		next.ServeHTTP(w, r)
	})
}

func (s HttpServer) Start() {
	startHandlers(s.Handlers)

	logger.Info(fmt.Sprintf("🚀 Server running on http://127.0.0.1:%d", s.Port))

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", s.Port),
		middlewares(http.DefaultServeMux, s.Middlewares),
	)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}
