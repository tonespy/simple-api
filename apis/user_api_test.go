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

func TestUserAPI_GetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1234", nil)
	assert.Nil(t, err)

	// This should fail because we didn't provide any user id
	recorder := mockRequestHandler(req, "GET", "/user/:id", getUser)
	assert.Equal(t, recorder.Code, http.StatusNotFound)

	// This should fail because an invalid UserID was provided
	req, err = http.NewRequest("GET", "/user/abc", nil)
	assert.Nil(t, err)

	recorder = mockRequestHandler(req, "GET", "/user/:id", getUser)
	assert.Contains(t, recorder.Body.String(), "Invalid ID")
	assert.Equal(t, recorder.Code, http.StatusNotFound)

	// Create user
	bufferBody := `{"first_name": "Abubakar", "last_name": "Oladeji", "email": "abc@email.com", "password": "password"}`
	req, err = http.NewRequest("POST", "/user", bytes.NewBufferString(bufferBody))
	assert.Nil(t, err)

	recorder = mockRequestHandler(req, "POST", "/user", createUser)
	assert.Equal(t, recorder.Code, http.StatusCreated)

	// This should fail because we provided an id that isn't in the record
	req, err = http.NewRequest("GET", "/user/1234", nil)
	assert.Nil(t, err)

	recorder = mockRequestHandler(req, "GET", "/user/:id", getUser)
	assert.Equal(t, recorder.Code, http.StatusNotFound)

	// This should pass with because, we provided a valid user ID
	req, err = http.NewRequest("GET", "/user/1", nil)
	assert.Nil(t, err)

	recorder = mockRequestHandler(req, "GET", "/user/:id", getUser)
	assert.Equal(t, recorder.Code, http.StatusFound)
}

// mockRequestHandler :- Mocks a handler and returns a httptest.ResponseRecorder
func mockRequestHandler(req *http.Request, method string, path string, reqHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, path, reqHandler)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	return recorder
}
