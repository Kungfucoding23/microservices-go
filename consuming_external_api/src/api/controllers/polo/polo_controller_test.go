package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/tests"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	c := tests.GetMockedContext(request, response)

	MarcoController(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
