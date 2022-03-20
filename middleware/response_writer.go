package middleware

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(rw http.ResponseWriter) *responseWriter {
	return &responseWriter{rw, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
