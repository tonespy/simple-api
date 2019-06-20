package apis

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tonespy/simple-api/models"
	"github.com/tonespy/simple-api/response"
	"github.com/tonespy/simple-api/router"
)

// ok represents types capable of validating
// themselves.
type ok interface {
	OK() error
}

// CreateUser :- Handler for creating a user
// post /user
func createUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.CreatedAt = time.Now().Local().String()

	resp := response.GenericResponse(201, "User Created Successfully.", user)

	response.WriteResponse(w, resp)
}

// GenerateUserRoutes :- Helper function for collating user routes
func GenerateUserRoutes() []router.Route {
	// Create user setup
	createUserRoute := router.Route{
		Name:            "Create User",
		Method:          "POST",
		Path:            "/user",
		HandlerFunction: createUser,
	}

	// collate all routes
	routes := []router.Route{createUserRoute}

	return routes
}

// decode can be this simple to start with, but can be extended
// later to support different formats and behaviours without
// changing the interface.
func decode(r *http.Request, v ok) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}
