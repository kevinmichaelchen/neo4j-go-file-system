package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type File struct {
	ResourceID uuid.UUID `json:"resourceID"`
	ParentID   uuid.UUID `json:"parentID"`
	Name       string    `json:"name"`
}

type FileService struct {
	DriverInfo DriverInfo
}

func (s *FileService) GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := uuid.Parse(idString)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	driver := GetDriver(s.DriverInfo)
	defer driver.Close()

	session := GetSession(driver)
	defer session.Close()

	file, err := getFileByID(session, id)

	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if file == nil {
		requestUtils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("No file found for ID: %s", idString))
		return
	}

	// TODO verify user can read file

	requestUtils.RespondWithJSON(w, http.StatusOK, file)
}

func getFileByID(session neo4j.Session, fileID uuid.UUID) (*File, error) {
	result, err := session.Run(`MATCH (f:File)<-[:CONTAINS_FILE]-(parent:Folder) WHERE f.resource_id = $resource_id RETURN f.resource_id, parent.resource_id, f.name`, map[string]interface{}{"resource_id": fileID.String()})
	if err != nil {
		return nil, err
	}
	// TODO should this code be safer (e.g., check for uuid parsing errors? check type casts?)
	if result.Next() {
		record := result.Record()
		return &File{
			ResourceID: uuid.Must(uuid.Parse(record.GetByIndex(0).(string))),
			ParentID:   uuid.Must(uuid.Parse(record.GetByIndex(1).(string))),
			Name:       record.GetByIndex(2).(string),
		}, nil
	}
	return nil, nil
}

func fileExists(session neo4j.Session, fileID uuid.UUID) (bool, error) {
	f, err := getFileByID(session, fileID)
	if err != nil {
		return false, err
	}
	return f != nil, nil
}
