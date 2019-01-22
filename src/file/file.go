package file

import (
	"log"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type File struct {
	ResourceID uuid.UUID `json:"resourceID"`
	ParentID   uuid.UUID `json:"parentID"`
	Name       string    `json:"name"`
}

type Controller struct {
	Service Service
}

type Service interface {
	GetFile(fileID uuid.UUID) (*File, *service.Error)
}

func (c *Controller) GetFileRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := uuid.Parse(idString)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, serviceError := c.Service.GetFile(id)
	if serviceError != nil {
		log.Println(serviceError.Error.Error())
		requestUtils.RespondWithError(w, serviceError.HttpCode, serviceError.ErrorMessage)
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, response)
}
