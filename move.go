package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type MoveService struct {
	DriverInfo DriverInfo
}

// MoveOperation represents a move operation of a file from one directory to another.
type MoveOperation struct {
	// SourceID is the resource ID of the file we wish to move.
	SourceID uuid.UUID `json:"source_id"`

	// DestinationID is the resource ID of the folder to which we wish to move the file.
	DestinationID uuid.UUID `json:"destination_id"`

	// NewName is the new name of the file, in the event the client wishes to rename the file.
	NewName string `json:"new_name"`
}

// Move moves a file
func (s *MoveService) Move(w http.ResponseWriter, r *http.Request) {
	var resource MoveOperation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := GetDriver(s.DriverInfo)
	defer driver.Close()

	session := GetSession(driver)
	defer session.Close()

	// TODO validate move operation, confirm src and dest exist and you have permission to read/write

	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Success"})
}
