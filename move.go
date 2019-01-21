package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type MoveService struct {
	DriverInfo DriverInfo
}

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

// Move moves a file
func (s *MoveService) Move(w http.ResponseWriter, r *http.Request) {
	var moveOperation MoveOperation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&moveOperation); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := GetDriver(s.DriverInfo)
	defer driver.Close()

	session := GetSession(driver)
	defer session.Close()

	source, err := getFileByID(session, moveOperation.SourceID)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if source == nil {
		requestUtils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("No file found for: %s", moveOperation.SourceID.String()))
		return
	}
	// TODO verify user can write to source file

	dest, err := getFolderByID(session, moveOperation.DestinationID)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if dest == nil {
		requestUtils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("No folder found for: %s", moveOperation.DestinationID.String()))
		return
	}
	// TODO verify user can write to destination folder

	if moveOperation.NewName != nil {
		// TODO validate new name is not too long; check if it's actually different
	}

	// Delete the old CONTAINS_FILE relationship
	tx, err := session.BeginTransaction()
	_, err = tx.Run(`MATCH (folder:Folder { resource_id: $old_parent_id })-[r:CONTAINS_FILE]->(file:File { resource_id: $file_id }) DELETE r`,
		map[string]interface{}{"file_id": source.ResourceID.String(), "old_parent_id": source.ParentID.String()})

	// Rollback if there's an error
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Create a new CONTAINS_FILE relationship
	_, err = tx.Run(`MATCH (folder:Folder), (file:File) WHERE folder.resource_id = $new_parent_id AND file.resource_id = $file_id CREATE (folder)-[r:CONTAINS_FILE]->(file) RETURN type(r)`,
		map[string]interface{}{"file_id": source.ResourceID.String(), "new_parent_id": dest.ResourceID.String()})

	// Rollback if there's an error
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Success"})
}
