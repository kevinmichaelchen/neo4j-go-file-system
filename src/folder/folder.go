package folder

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
)

type FolderInput struct {
	ResourceID string  `json:"resourceID"`
	ParentID   *string `json:"parentID"`
	Name       string  `json:"name"`
	RevisionID string  `json:"revisionID"`
}

type Folder struct {
	ResourceID uuid.UUID  `json:"resourceID"`
	ParentID   *uuid.UUID `json:"parentID"`
	Name       string     `json:"name"`
	RevisionID uuid.UUID  `json:"revisionID"`
}

type Controller struct {
	Service Service
}

type Service interface {
	CreateFolder(ctx context.Context, folder FolderInput) (*Folder, *service.Error)
	GetFolder(ctx context.Context, folder FolderInput) (*Folder, *service.Error)
	UpdateFolder(ctx context.Context, folder FolderInput) (*Folder, *service.Error)
	DeleteFolder(ctx context.Context, folder FolderInput) (*Folder, *service.Error)
}
