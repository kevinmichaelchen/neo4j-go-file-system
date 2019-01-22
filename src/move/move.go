package move

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type Service struct {
	DriverInfo neo.DriverInfo
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

// MoveRequestHandler moves a file
func (s *Service) MoveRequestHandler(w http.ResponseWriter, r *http.Request) {
	var moveOperation MoveOperation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&moveOperation); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	source, err := file.GetFileByID(session, moveOperation.SourceID)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if source == nil {
		requestUtils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("No file found for: %s", moveOperation.SourceID.String()))
		return
	}
	// TODO verify user can write to source file

	dest, err := folder.GetFolderByID(session, moveOperation.DestinationID)
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
		// TODO validate new name is not too long
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

	// Update the file's name if it was changed
	if moveOperation.NewName != nil {
		_, err = tx.Run(`MATCH (f:File {resource_id: $resource_id}) SET f.name = $name RETURN f.name`,
			map[string]interface{}{"resource_id": source.ResourceID.String(), "name": *moveOperation.NewName})

		// Rollback if there's an error
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Success"})
}
