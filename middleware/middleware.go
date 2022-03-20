package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Middleware struct {
	l *zap.SugaredLogger
}

func NewMiddleware(l *zap.SugaredLogger) *Middleware {
	return &Middleware{l: l}
}

func (m *Middleware) HTTPRouter(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		crw := newResponseWriter(rw)
		next(crw, req, ps)
		m.requestLog(req, crw.statusCode)
	}
}

func (m *Middleware) HTTPStandardHandler(next http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		crw := newResponseWriter(rw)
		next.ServeHTTP(crw, req)
		m.requestLog(req, crw.statusCode)
	}
}

func (m *Middleware) requestLog(r *http.Request, code int) {
	m.l.With(
		"type", "access",
		"event", "request",
		"remote-ip", strings.Split(r.RemoteAddr, ":")[0],
		"host", r.Host,
		"url-path", r.RequestURI,
		"method", r.Method,
		"status-code", code,
	).Info(
		fmt.Sprintf("%d -> %s %s", code, r.Method, r.RequestURI),
	)
}
