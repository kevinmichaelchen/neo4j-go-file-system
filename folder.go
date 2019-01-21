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

func getFolderByID(session neo4j.Session, folderID uuid.UUID) (*Folder, error) {
	result, err := session.Run(`MATCH (child:Folder { resource_id: $resource_id }) OPTIONAL MATCH (child)<-[:CONTAINS_FOLDER]-(parent:Folder) RETURN child.resource_id, parent.resource_id, child.name`, map[string]interface{}{"resource_id": folderID.String()})
	if err != nil {
		return nil, err
	}
	if result.Next() {
		record := result.Record()
		f := &Folder{
			ResourceID: uuid.Must(uuid.Parse(record.GetByIndex(0).(string))),
			Name:       record.GetByIndex(2).(string),
		}
		parentIDString := record.GetByIndex(1)

		// A folder might not have a parent, so we do a nil check to avoid type assertion errors
		if parentIDString != nil {
			parentID, err := uuid.Parse(parentIDString.(string))
			if err == nil {
				f.ParentID = &parentID
			}
		}

		return f, nil
	}
	return nil, nil
}

func folderExists(session neo4j.Session, folderID uuid.UUID) (bool, error) {
	f, err := getFolderByID(session, folderID)
	if err != nil {
		return false, err
	}
	return f != nil, nil
}
