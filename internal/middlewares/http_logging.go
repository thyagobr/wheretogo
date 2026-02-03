package middlewares

import (
		"log"
		"net/http"
		"time"
)

func HttpLoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        next.ServeHTTP(w, r)

        duration := time.Since(start)

        log.Printf(
            "[%s] %s %s params=%v (%s)",
            start.Format(time.RFC3339),
            r.Method,
            r.URL.Path,
            r.URL.Query(),
            duration,
        )
    })
}
