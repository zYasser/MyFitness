package middleware

import (
	"context"
	"net/http"

	"github.com/zYasser/MyFitness/utils"
)

type ctxKeyLogger struct{}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := utils.GetLogger()
		newCtx := context.WithValue(r.Context(), ctxKeyLogger{}, logger)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})

}

func FromContext(ctx context.Context) *utils.Logger {
	logger := ctx.Value(ctxKeyLogger{}).(*utils.Logger)
	return logger
}
