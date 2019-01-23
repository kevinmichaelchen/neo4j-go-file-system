package neo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

type Service struct {
	DriverInfo neo.DriverInfo
}

func NewService(driverInfo neo.DriverInfo) *Service {
	return &Service{DriverInfo: driverInfo}
}

func (s *Service) CreateFile(context context.Context, in file.File) (*file.File, *service.Error) {
	return nil, service.Unimplemented()
}

func (s *Service) GetFile(context context.Context, in file.File) (*file.File, *service.Error) {
	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	resource, err := GetFileByID(session, in.ResourceID)

	userID, svcErr := service.GetUserID(context)
	if err != nil {
		return nil, svcErr
	}

	log.Printf("Got USER ID %d\n", userID)

	if err != nil {
		return nil, service.Internal(err)
	}

	if resource == nil {
		return nil, service.NotFound(fmt.Sprintf("No file found for ID: %s", in.ResourceID.String()))
	}

	// TODO verify user can read file

	return resource, nil
}

func (s *Service) UpdateFile(context context.Context, in file.File) (*file.File, *service.Error) {
	return nil, service.Unimplemented()
}

func (s *Service) DeleteFile(context context.Context, in file.File) (*file.File, *service.Error) {
	return nil, service.Unimplemented()
}

func GetFileByID(session neo4j.Session, fileID uuid.UUID) (*file.File, error) {
	result, err := session.Run(`MATCH (f:File)<-[:CONTAINS_FILE]-(parent:Folder) WHERE f.resource_id = $resource_id RETURN f.resource_id, parent.resource_id, f.name`, map[string]interface{}{"resource_id": fileID.String()})
	if err != nil {
		return nil, err
	}
	// TODO should this code be safer? (e.g., check for uuid parsing errors? check type casts?)
	if result.Next() {
		record := result.Record()
		return &file.File{
			ResourceID: uuid.Must(uuid.Parse(record.GetByIndex(0).(string))),
			ParentID:   uuid.Must(uuid.Parse(record.GetByIndex(1).(string))),
			Name:       record.GetByIndex(2).(string),
		}, nil
	}
	return nil, nil
}

func fileExists(session neo4j.Session, fileID uuid.UUID) (bool, error) {
	f, err := GetFileByID(session, fileID)
	if err != nil {
		return false, err
	}
	return f != nil, nil
}
