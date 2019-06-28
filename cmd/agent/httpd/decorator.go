package httpd

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type LoggerWrapperResponseWriter struct {
	w             http.ResponseWriter
	logger        *zap.Logger
	statusCode    int
	contentLength int
}

func (lr *LoggerWrapperResponseWriter) Header() http.Header {
	return lr.w.Header()
}

func (lr *LoggerWrapperResponseWriter) Write(content []byte) (int, error) {
	length, err := lr.w.Write(content)
	lr.contentLength = lr.contentLength + length
	return length, err
}

func (lr *LoggerWrapperResponseWriter) WriteHeader(statusCode int) {
	lr.statusCode = statusCode
	lr.w.WriteHeader(statusCode)
}

var _ http.ResponseWriter = &LoggerWrapperResponseWriter{}

type Decorator struct {
	next   http.Handler
	logger *zap.Logger
}

func (d *Decorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	loggerWriter := &LoggerWrapperResponseWriter{
		w:      w,
		logger: d.logger,
	}

	start := time.Now()
	d.next.ServeHTTP(loggerWriter, r)
	end := time.Now()
	duration := end.Sub(start)

	httpAction.With(prometheus.Labels{
		"path":        r.URL.Path,
		"method":      r.Method,
		"status_code": strconv.Itoa(loggerWriter.statusCode),
	}).Inc()

	d.logger.Info("",
		zap.String("remote", r.RemoteAddr),
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("query", r.URL.RawQuery),
		zap.Int("code", loggerWriter.statusCode),
		zap.Int("size", loggerWriter.contentLength),
		zap.Duration("latency", duration))
}

var _ http.Handler = &Decorator{}

func NewHttpHandler(hdr http.Handler, logger *zap.Logger) http.Handler {
	return &Decorator{
		next:   hdr,
		logger: logger,
	}
}
