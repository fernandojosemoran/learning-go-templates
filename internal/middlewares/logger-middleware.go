package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fernandojosemoran/go-templates/pkg/logger"
)

func LoggerMiddleware(response http.ResponseWriter, request *http.Request) {
	var pathCompleted string = request.URL.Path + request.URL.RawQuery
	start := time.Now()
	logger.Info(fmt.Sprintf("%s %s - %v\n", request.Method, pathCompleted, time.Since(start)))
}
