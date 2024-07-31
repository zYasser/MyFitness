package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
)

type ctxKeyLogger struct{}

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
var logger =&Logger{InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime | log.Lshortfile ), ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        newCtx := context.WithValue(r.Context(), ctxKeyLogger{}, logger)
        next.ServeHTTP(w, r.WithContext(newCtx))
    })
}

func FromContext(ctx context.Context) (*Logger) {
    logger := ctx.Value(ctxKeyLogger{}).(*Logger)
    return logger
}
