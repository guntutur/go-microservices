package ping

import (
	"github.com/guntutur/go-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "pong", pong)
}

func TestPong(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	c := test_utils.GetMockContext(request, response)

	Pong(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "pong", response.Body.String())
}
