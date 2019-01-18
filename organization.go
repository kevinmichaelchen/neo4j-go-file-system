package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Organization struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func CreateOrganization(session neo4j.Session, organization Organization) error {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			`CREATE (Organization {resource_id: $id, name: $name})`, orgToMap(organization))
	})
	if err != nil {
		return err
	}
	return nil
}

func organizationExists(session neo4j.Session, organization Organization) (bool, error) {
	res, err := session.Run(`MATCH (o:Organization {name: $name}) RETURN o.name`, map[string]interface{}{"name": organization.Name})
	if err != nil {
		return false, err
	}
	if res.Next() {
		e := res.Record().GetByIndex(0).(string)
		return e != "", nil
	}
	return false, nil
}

func orgToMap(organization Organization) map[string]interface{} {
	return map[string]interface{}{
		"id":   organization.ID.String(),
		"name": organization.Name,
	}
}

type OrganizationService struct {
	DriverInfo DriverInfo
}

// CreateOrganization creates an org
func (s *OrganizationService) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var resource Organization
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := GetDriver(s.DriverInfo)
	defer driver.Close()

	session := GetSession(driver)
	defer session.Close()

	// Set the ID
	resource.ID = uuid.Must(uuid.NewRandom())

	// TODO validate org resource

	exists, err := organizationExists(session, resource)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if exists {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Org already exists with that name")
		return
	}

	err = CreateOrganization(session, resource)

	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Created org"})
}
