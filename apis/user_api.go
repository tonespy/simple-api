package apis

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/tonespy/simple-api/errors"
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
	err := decode(r, &user)
	if err != nil {
		validationData := errors.Params{
			"first_name": "required",
			"last_name":  "required",
			"password":   "required",
			"email":      "required",
		}
		errResp := errors.NewAPIError(http.StatusBadRequest, "BAD_REQUEST", "Please provide valid user data.", validationData)
		errors.WriteErrorResponse(w, errResp)
	}

	user.CreatedAt = time.Now().Local().String()
	user.UpdatedAt = time.Now().Local().String()

	models.UserStore[strconv.Itoa(user.ID)] = user

	resp := response.GenericResponse(http.StatusCreated, "User Created Successfully.", user)

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
