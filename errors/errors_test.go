package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {

	e := InternalServerError(errors.New(""))
	assert.Equal(t, http.StatusInternalServerError, e.Status)
}

func TestNotFound(t *testing.T) {

	e := NotFound("/xyz")
	assert.Equal(t, http.StatusNotFound, e.Status)
}

func TestUnauthorized(t *testing.T) {

	e := Unauthorized("Invalid email or password provided")
	assert.Equal(t, http.StatusUnauthorized, e.Status)
}

func TestGenericError(t *testing.T) {

	e := GenericError(450, Params{"x": "xyz"},  "INVALID_TOKEN", "Provide a valid user token")
	assert.Equal(t, 450, e.Status)
	assert.Equal(t, "INVALID_TOKEN", e.ErrorCode)
}
