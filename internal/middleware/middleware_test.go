package middleware

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMiddleware(t *testing.T) {
	var buf bytes.Buffer

	logger := slog.New(slog.NewTextHandler(&buf, nil))

	middleware := NewMiddleware(logger)
	assert.NotNil(t, middleware.log)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	rr := httptest.NewRecorder()

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}

	loggingHandler := middleware.Logging(testHandler)

	loggingHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "OK", rr.Body.String())

	logOutput := buf.String()
	assert.Contains(t, logOutput, "started handling request")
	assert.Contains(t, logOutput, "sent response to request")

}
