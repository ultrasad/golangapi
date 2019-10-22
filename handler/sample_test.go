package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	e := initEcho()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	// Using the ServerHTTP on echo will trigger the router and middleware
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, rec.Body.String(), "body output")
	assert.Equal(t, rec.Header().Get(echo.HeaderServer), "Echo/3.0")
}
