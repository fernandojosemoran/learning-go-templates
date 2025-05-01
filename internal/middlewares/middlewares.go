package middlewares

import (
	"net/http"
)

var Middlewares []http.HandlerFunc = []http.HandlerFunc{
	LoggerMiddleware,
}
