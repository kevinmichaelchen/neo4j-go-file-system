package main

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type File struct {
	ResourceID uuid.UUID `json:"resourceID"`
	Name       string    `json:"name"`
}

func fileExists(session neo4j.Session, uuid uuid.UUID) (bool, error) {
	res, err := session.Run(`MATCH (f:File {resource_id: $resource_id}) RETURN f.name`, map[string]interface{}{"resource_id": uuid.String()})
	if err != nil {
		return false, err
	}
	if res.Next() {
		e := res.Record().GetByIndex(0).(string)
		return e != "", nil
	}
	return false, nil
}
