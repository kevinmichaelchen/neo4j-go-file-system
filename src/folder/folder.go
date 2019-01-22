package folder

import (
	"github.com/google/uuid"
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
}
