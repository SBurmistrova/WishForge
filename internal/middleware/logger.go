package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func CreateLogger() (*slog.Logger, error) {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf(
		"logs/%s.log",
		time.Now().Format("2006-01-02"),
	)

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(file, nil))

	return logger, nil
}

func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			next.ServeHTTP(rw, r)

			duration := time.Since(start)

			logger.Info("request completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rw.status),
				slog.Duration("duration", duration),
				slog.String("ip", r.RemoteAddr),
			)
		})
	}
}
