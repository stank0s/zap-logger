package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponseWriter(t *testing.T) {
	// when: creating new response writer
	r := newResponseWriter(httptest.NewRecorder())

	// then: expected status code has been set
	assert.Equal(t, http.StatusOK, r.statusCode)
}

func TestWriteHeader(t *testing.T) {
	// given: test subject
	rw := newResponseWriter(httptest.NewRecorder())

	// when: setting status code to 400
	rw.statusCode = http.StatusBadRequest

	// then: expected value has been set
	assert.Equal(t, http.StatusBadRequest, rw.statusCode)

	// when: setting status code to 500
	rw.WriteHeader(http.StatusInternalServerError)

	// then: expected value has been set
	assert.Equal(t, http.StatusInternalServerError, rw.statusCode)
}
