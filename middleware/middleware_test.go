package middleware

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func TestNewMiddleware(t *testing.T) {
	// given: test logger
	l := zaptest.NewLogger(t).Sugar()

	// when: creating new middleware
	m := NewMiddleware(l)

	// then: object has correct type
	assert.IsType(t, &Middleware{}, m)
}

func TestHTTPRouter(t *testing.T) {
	// given: test logger mock
	var buffer bytes.Buffer

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buffer)

	l := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)).
		Sugar()

	// and: test subject
	mdw := NewMiddleware(l)

	// and: mocked middleware
	m := mdw.HTTPRouter(httprouter.Handle(func(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rw.WriteHeader(http.StatusOK)
	}))

	// and: test responseWriter
	rw := httptest.NewRecorder()

	// and: test request
	req, err := http.NewRequest(http.MethodGet, "http://test.com", nil)
	assert.NoError(t, err)
	req.RequestURI = "/"

	// when: middleware function called
	m(rw, req, httprouter.Params{})
	writer.Flush()

	// then: expected results returned
	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, buffer.String(), `200 -> GET /`)
	assert.Contains(t, buffer.String(), `"type": "access"`)
	assert.Contains(t, buffer.String(), `"event": "request"`)
	assert.Contains(t, buffer.String(), `"method": "GET"`)
	assert.Contains(t, buffer.String(), `"status-code": 200`)
	assert.Contains(t, buffer.String(), `"url-path": "/"`)
}

func TestHTTPStandardHandler(t *testing.T) {
	// given: test logger mock
	var buffer bytes.Buffer

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buffer)

	l := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)).
		Sugar()

	// and: test subject
	mdw := NewMiddleware(l)

	// and: mocked middleware
	m := mdw.HTTPStandardHandler(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))

	// and: test responseWriter
	rw := httptest.NewRecorder()

	// and: test request
	req, err := http.NewRequest(http.MethodDelete, "http://test.com", nil)
	assert.NoError(t, err)
	req.RequestURI = "/someurl"

	// when: middleware function called
	m(rw, req)
	writer.Flush()

	// then: expected results returned
	assert.Equal(t, http.StatusInternalServerError, rw.Code)
	assert.Contains(t, buffer.String(), `500 -> DELETE /someurl`)
	assert.Contains(t, buffer.String(), `"type": "access"`)
	assert.Contains(t, buffer.String(), `"event": "request"`)
	assert.Contains(t, buffer.String(), `"method": "DELETE"`)
	assert.Contains(t, buffer.String(), `"status-code": 500`)
	assert.Contains(t, buffer.String(), `"url-path": "/someurl"`)
}
