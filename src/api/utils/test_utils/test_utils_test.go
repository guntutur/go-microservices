package test_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMockContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:1234/sumtin", nil)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock" : {"true"}}
	c := GetMockContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "1234", c.Request.URL.Port())
	assert.EqualValues(t, "/sumtin", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}
