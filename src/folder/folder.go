package folder

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
)

type Folder struct {
	ResourceID uuid.UUID  `json:"resourceID"`
	ParentID   *uuid.UUID `json:"parentID"`
	Name       string     `json:"name"`
}

type Controller struct {
	Service Service
}

type Service interface {
	CreateFolder(ctx context.Context, folder Folder) (*Folder, *service.Error)
	GetFolder(ctx context.Context, folder Folder) (*Folder, *service.Error)
	UpdateFolder(ctx context.Context, folder Folder) (*Folder, *service.Error)
	DeleteFolder(ctx context.Context, folder Folder) (*Folder, *service.Error)
}
