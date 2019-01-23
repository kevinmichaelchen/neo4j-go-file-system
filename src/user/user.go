package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/service"

	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type User struct {
	ResourceID   int64  `json:"resourceID"`
	EmailAddress string `json:"emailAddress"`
	FullName     string `json:"fullName"`
}

type Controller struct {
	Service Service
}

type Service interface {
	CreateUser(user User) (*User, *service.Error)
	GetUser(user User) (*User, *service.Error)
	UpdateUser(user User) (*User, *service.Error)
	DeleteUser(user User) (*User, *service.Error)
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
