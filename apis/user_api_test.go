package apis

import (
	"fmt"
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

func TestUserAPI_CreateUser(t *testing.T) {
	req, err := http.NewRequest("POST", "/user", nil)
	assert.Nil(t, err, "Invalid Post request")

	createTest1 := mockRequestHandler(req, "POST", "/user", createUser)
	fmt.Println("Create Test Code: ", createTest1.Code)
}

// mockRequestHandler :- Mocks a handler and returns a httptest.ResponseRecorder
func mockRequestHandler(req *http.Request, method string, path string, reqHandler func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, path, reqHandler)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	return recorder
}
