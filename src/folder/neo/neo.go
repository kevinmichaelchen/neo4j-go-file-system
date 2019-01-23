package neo

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Service struct {
	DriverInfo neo.DriverInfo
}

func NewService(driverInfo neo.DriverInfo) *Service {
	return &Service{DriverInfo: driverInfo}
}

func (s *Service) CreateFolder(ctx context.Context, folder folder.FolderInput) (*folder.Folder, *service.Error) {
	if folder.ParentID != nil {
		_, err := uuid.Parse(*folder.ParentID)
		if err != nil {
			return nil, service.InvalidArgument("invalid folder ID")
		}
	}

	return nil, service.Unimplemented()
}

func (s *Service) GetFolder(ctx context.Context, folder folder.FolderInput) (*folder.Folder, *service.Error) {
	_, err := uuid.Parse(folder.ResourceID)
	if err != nil {
		return nil, service.InvalidArgument("invalid folder ID")
	}

	return nil, service.Unimplemented()
}

func (s *Service) UpdateFolder(ctx context.Context, in folder.FolderInput) (*folder.Folder, *service.Error) {
	if strings.TrimSpace(in.ResourceID) == "" {
		return nil, service.InvalidArgument("must provide folder ID")
	}

	_, err := uuid.Parse(in.ResourceID)
	if err != nil {
		return nil, service.InvalidArgument("invalid folder ID")
	}

	if in.ParentID != nil {
		_, err := uuid.Parse(*in.ParentID)
		if err != nil {
			return nil, service.InvalidArgument("invalid parent ID")
		}
	}

	return nil, service.Unimplemented()
}

func (s *Service) DeleteFolder(ctx context.Context, folder folder.FolderInput) (*folder.Folder, *service.Error) {
	_, err := uuid.Parse(folder.ResourceID)
	if err != nil {
		return nil, service.InvalidArgument("invalid folder ID")
	}
	return nil, service.Unimplemented()
}

func GetFolderByID(session neo4j.Session, folderID uuid.UUID) (*folder.Folder, error) {
	result, err := session.Run(`MATCH (child:Folder { resource_id: $resource_id }) OPTIONAL MATCH (child)<-[:CONTAINS_FOLDER]-(parent:Folder) RETURN child.resource_id, parent.resource_id, child.name`, map[string]interface{}{"resource_id": folderID.String()})
	if err != nil {
		return nil, err
	}
	if result.Next() {
		record := result.Record()
		f := &folder.Folder{
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
	f, err := GetFolderByID(session, folderID)
	if err != nil {
		return false, err
	}
	return f != nil, nil
}
