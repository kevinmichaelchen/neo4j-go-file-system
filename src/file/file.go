package file

import (
	"context"
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
	CreateFile(ctx context.Context, file File) (*File, *service.Error)
	GetFile(ctx context.Context, file File) (*File, *service.Error)
	UpdateFile(ctx context.Context, file File) (*File, *service.Error)
	DeleteFile(ctx context.Context, file File) (*File, *service.Error)
}

func (c *Controller) GetFileRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileIDString := vars["id"]
	fileID, err := uuid.Parse(fileIDString)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, serviceError := c.Service.GetFile(service.CreateUserContext(11), File{ResourceID: fileID})
	if serviceError != nil {
		log.Println(serviceError.Error.Error())
		requestUtils.RespondWithError(w, serviceError.HttpCode, serviceError.ErrorMessage)
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, response)
}
