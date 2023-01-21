package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleOriginUrl(t *testing.T) {
	s := NewServer(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/origin_url", nil)
	s.handleOriginUrl().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "origin string")
}
