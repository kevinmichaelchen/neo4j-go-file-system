package main

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Folder struct {
	ResourceID uuid.UUID  `json:"resourceID"`
	ParentID   *uuid.UUID `json:"parentID"`
	Name       string     `json:"name"`
}

type FolderService struct {
	DriverInfo DriverInfo
}

func folderExists(session neo4j.Session, folderID uuid.UUID) (bool, error) {
	res, err := session.Run(`MATCH (f:Folder {resource_id: $resource_id}) RETURN f.name`, map[string]interface{}{"resource_id": folderID.String()})
	if err != nil {
		return false, err
	}
	if res.Next() {
		e := res.Record().GetByIndex(0).(string)
		return e != "", nil
	}
	return false, nil
}
