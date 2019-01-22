package move

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/service"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

// MoveOperation represents a move operation of a file from one directory to another.
type MoveOperation struct {
	// SourceID is the resource ID of the file we wish to move.
	SourceID uuid.UUID `json:"sourceID"`

	// DestinationID is the resource ID of the folder to which we wish to move the file.
	DestinationID uuid.UUID `json:"destinationID"`

	// NewName is the new name of the file, in the event the client wishes to rename the file.
	// We use a pointer to indicate that this field is optional, in which case no rename occurs.
	NewName *string `json:"newName"`
}

type Controller struct {
	Service Service
}

type Service interface {
	Move(operation MoveOperation) (*MoveOperation, *service.Error)
}

// MoveRequestHandler moves a file
func (c *Controller) MoveRequestHandler(w http.ResponseWriter, r *http.Request) {
	var resource MoveOperation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	response, serviceError := c.Service.Move(resource)
	if serviceError != nil {
		log.Println(serviceError.Error.Error())
		requestUtils.RespondWithError(w, serviceError.HttpCode, serviceError.ErrorMessage)
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, response)
}
