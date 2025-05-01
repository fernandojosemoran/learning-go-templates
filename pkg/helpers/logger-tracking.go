package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fernandojosemoran/go-templates/pkg/logger"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("host: %s\n\n, method: %s\n\n RawQuery: %s\n\n, RawPath: %s\n\n", r.Host, r.Method, r.URL.RawQuery)
		start := time.Now()
		logger.Info(fmt.Sprintf("%s %s - %v\n", r.Method, r.URL.Path+r.URL.RawQuery, time.Since(start)))
		next.ServeHTTP(w, r)
	})
}
