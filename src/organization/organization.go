package organization

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/service"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type Organization struct {
	ResourceID uuid.UUID `json:"resourceID"`
	Name       string    `json:"name"`
}

type Controller struct {
	Service Service
}

type Service interface {
	CreateOrganization(organization Organization) (*Organization, *service.Error)
}

// CreateOrganizationRequestHandler creates an org
func (c *Controller) CreateOrganizationRequestHandler(w http.ResponseWriter, r *http.Request) {
	var resource Organization
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	response, serviceError := c.Service.CreateOrganization(resource)
	if serviceError != nil {
		log.Println(serviceError.Error.Error())
		requestUtils.RespondWithError(w, serviceError.HttpCode, serviceError.ErrorMessage)
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, response)

}
