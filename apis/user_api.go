package apis

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	appError "github.com/tonespy/simple-api/errors"
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
// POST /user
func createUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var user models.User
	err := decode(r, &user)
	if err != nil {
		validationData := appError.Params{
			"first_name": "required",
			"last_name":  "required",
			"password":   "required",
			"email":      "required",
		}
		errResp := appError.NewAPIError(http.StatusBadRequest, "BAD_REQUEST", "Please provide valid user data.", validationData)
		appError.WriteErrorResponse(w, errResp)
		return
	}

	user.CreatedAt = time.Now().Local().String()
	user.UpdatedAt = time.Now().Local().String()
	user.ID = models.GenerateUserID()

	models.UserStore[strconv.Itoa(user.ID)] = user

	resp := response.GenericResponse(http.StatusCreated, "User Created Successfully.", user)

	response.WriteResponse(w, resp)
}

// getUser :- Handler for getting user information
// GET /user/:id
func getUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := params.ByName("id")
	if _, err := strconv.Atoi(userID); err != nil {
		apiError := appError.NotFound("Invalid ID " + userID)
		appError.WriteErrorResponse(w, apiError)
		return
	}

	if userInfo, ok := models.UserStore[userID]; ok {
		resp := response.GenericResponse(http.StatusFound, "User found successfully", userInfo)
		response.WriteResponse(w, resp)
		return
	}

	apiError := appError.NotFound("Invalid ID " + userID)
	appError.WriteErrorResponse(w, apiError)
}

// updateUser :- Handler for updating user information
// PUT /user/:id
func updateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := params.ByName("id")
	if _, err := strconv.Atoi(userID); err != nil {
		apiError := appError.NotFound("Invalid ID " + userID)
		appError.WriteErrorResponse(w, apiError)
		return
	}

	if _, ok := models.UserStore[userID]; !ok {
		apiError := appError.NotFound("Invalid ID " + userID)
		appError.WriteErrorResponse(w, apiError)
		return
	}

	userInfo := models.UserStore[userID]

	var updatedUser models.User
	errorParam := appError.Params{"allowedParams": []string{"first_name", "last_name"}}

	if r.Body == nil {
		appError.WriteErrorResponse(w, appError.GenericError(http.StatusBadRequest, errorParam, "INVALID_DATA", "Please provide valid data"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		appError.WriteErrorResponse(w, appError.GenericError(http.StatusBadRequest, errorParam, "INVALID_DATA", "Please provide valid data"))
		return
	}

	if len(updatedUser.Firstname) > 0 {
		userInfo.Firstname = updatedUser.Firstname
	}

	if len(updatedUser.Lastname) > 0 {
		userInfo.Lastname = updatedUser.Lastname
	}

	userInfo.UpdatedAt = time.Now().Local().String()

	models.UserStore[userID] = userInfo

	resp := response.GenericResponse(http.StatusOK, "User updated successfully.", userInfo)
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

	// Get user setup
	getUserRoute := router.Route{
		Name:            "Get User",
		Method:          "GET",
		Path:            "/user/:id",
		HandlerFunction: getUser,
	}

	// Update user setup
	updateUserRoute := router.Route{
		Name:            "Update User",
		Method:          "PUT",
		Path:            "/user/:id",
		HandlerFunction: updateUser,
	}

	// collate all routes
	routes := []router.Route{createUserRoute, getUserRoute, updateUserRoute}

	return routes
}

// decode can be this simple to start with, but can be extended
// later to support different formats and behaviours without
// changing the interface.
func decode(r *http.Request, v ok) error {
	if r.Body == nil {
		return errors.New("Invalid Body")
	}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}
