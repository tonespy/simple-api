package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_GenericResponse(t *testing.T) {

	resp := GenericResponse(http.StatusCreated, "User created successfully", nil)
	assert.Equal(t, resp.Data, nil)
	assert.Equal(t, resp.StatusCode(), http.StatusCreated)
	assert.Equal(t, resp.Status, http.StatusCreated)
}
