package neo

import (
	"fmt"
	"net/http"

	fileNeo "github.com/kevinmichaelchen/neo4j-go-file-system/file/neo"
	folderNeo "github.com/kevinmichaelchen/neo4j-go-file-system/folder/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/move"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
)

type Service struct {
	DriverInfo neo.DriverInfo
}

func NewService(driverInfo neo.DriverInfo) *Service {
	return &Service{DriverInfo: driverInfo}
}

func (s *Service) Move(resource move.MoveOperation) (*move.MoveOperation, *service.Error) {
	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	source, err := fileNeo.GetFileByID(session, resource.SourceID)
	if err != nil {
		return nil, service.NewError(http.StatusInternalServerError, err.Error(), err)
	}
	if source == nil {
		return nil, service.NewError(http.StatusNotFound, fmt.Sprintf("No file found for: %s", resource.SourceID.String()), nil)
	}
	// TODO verify user can write to source file

	dest, err := folderNeo.GetFolderByID(session, resource.DestinationID)
	if err != nil {
		return nil, service.NewError(http.StatusInternalServerError, err.Error(), err)
	}
	if dest == nil {
		return nil, service.NewError(http.StatusNotFound, fmt.Sprintf("No folder found for: %s", resource.DestinationID.String()), nil)
	}
	// TODO verify user can write to destination folder

	if resource.NewName != nil {
		// TODO validate new name is not too long
	}

	// Delete the old CONTAINS_FILE relationship
	tx, err := session.BeginTransaction()
	_, err = tx.Run(`MATCH (folder:Folder { resource_id: $old_parent_id })-[r:CONTAINS_FILE]->(file:File { resource_id: $file_id }) DELETE r`,
		map[string]interface{}{"file_id": source.ResourceID.String(), "old_parent_id": source.ParentID.String()})

	// Rollback if there's an error
	if svcError := neo.RollbackIfError(tx, err); svcError != nil {
		return nil, svcError
	}

	// Create a new CONTAINS_FILE relationship
	_, err = tx.Run(`MATCH (folder:Folder), (file:File) WHERE folder.resource_id = $new_parent_id AND file.resource_id = $file_id CREATE (folder)-[r:CONTAINS_FILE]->(file) RETURN type(r)`,
		map[string]interface{}{"file_id": source.ResourceID.String(), "new_parent_id": dest.ResourceID.String()})

	// Rollback if there's an error
	if svcError := neo.RollbackIfError(tx, err); svcError != nil {
		return nil, svcError
	}

	// Update the file's name if it was changed
	if resource.NewName != nil {
		_, err = tx.Run(`MATCH (f:File {resource_id: $resource_id}) SET f.name = $name RETURN f.name`,
			map[string]interface{}{"resource_id": source.ResourceID.String(), "name": *resource.NewName})

		// Rollback if there's an error
		if svcError := neo.RollbackIfError(tx, err); svcError != nil {
			return nil, svcError
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, service.NewError(http.StatusInternalServerError, err.Error(), err)
	}

	return &resource, nil
}
