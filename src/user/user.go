package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type User struct {
	ResourceID   uuid.UUID `json:"resourceID"`
	EmailAddress string    `json:"emailAddress"`
	FullName     string    `json:"fullName"`
}

type Controller struct {
	Service Service
}

type Service interface {
	CreateUser(user User) (*User, *ServiceError)
}

type ServiceError struct {
	HttpCode     int
	ErrorMessage string
	Error        error
}

// CreateUserRequestHandler creates a user
func (c *Controller) CreateUserRequestHandler(w http.ResponseWriter, r *http.Request) {
	var resource User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	response, serviceError := c.Service.CreateUser(resource)
	if serviceError != nil {
		log.Println(serviceError.Error.Error())
		requestUtils.RespondWithError(w, serviceError.HttpCode, serviceError.ErrorMessage)
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, response)
}
