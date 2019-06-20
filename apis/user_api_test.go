package apis

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUserAPI_GenerateUserRoutes(t *testing.T) {
	routes := GenerateUserRoutes()
	assert.True(t, len(routes) > 0)
}

func TestUserAPI_FailingDecode(t *testing.T) {
	bufferBody := `{"email: "abc@email.com", "password": "password"}`
	req, err := http.NewRequest("POST", "/user", bytes.NewBufferString(bufferBody))
	assert.Nil(t, err)

	recorder := mockRequestHandler(req, "POST", "/user", createUser)
	assert.Equal(t, recorder.Code, http.StatusBadRequest)
}

func TestUserAPI_CreateUser(t *testing.T) {
	// This request should panic in the recorder because, we passed a nil body
	req, err := http.NewRequest("POST", "/user", nil)
	assert.Nil(t, err, "Invalid Post request")

	assert.Panics(t, func() { mockRequestHandler(req, "POST", "/user", createUser) })

	bufferBody := `{"email": "abc@email.com", "password": "password"}`
	req, err = http.NewRequest("POST", "/user", bytes.NewBufferString(bufferBody))
	assert.Nil(t, err)

	recorder := mockRequestHandler(req, "POST", "/user", createUser)
	assert.Equal(t, recorder.Code, http.StatusBadRequest)

	bufferBody = `{"first_name": "Abubakar", "last_name": "Oladeji", "email": "abc@email.com", "password": "password"}`
	req, err = http.NewRequest("POST", "/user", bytes.NewBufferString(bufferBody))
	assert.Nil(t, err)

	recorder = mockRequestHandler(req, "POST", "/user", createUser)
	assert.Equal(t, recorder.Code, http.StatusCreated)
}

// mockRequestHandler :- Mocks a handler and returns a httptest.ResponseRecorder
func mockRequestHandler(req *http.Request, method string, path string, reqHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, path, reqHandler)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	return recorder
}
