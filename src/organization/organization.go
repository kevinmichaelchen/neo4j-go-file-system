package organization

import (
	"encoding/json"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Organization struct {
	ResourceID uuid.UUID `json:"resourceID"`
	Name       string    `json:"name"`
}

func CreateOrganization(session neo4j.Session, organization Organization) error {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			`CREATE (Organization {resource_id: $resource_id, name: $name})`, orgToMap(organization))
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
		"resource_id": organization.ResourceID.String(),
		"name":        organization.Name,
	}
}

type Controller struct {
	DriverInfo neo.DriverInfo
}

// CreateOrganizationRequestHandler creates an org
func (c *Controller) CreateOrganizationRequestHandler(w http.ResponseWriter, r *http.Request) {
	var resource Organization
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := neo.GetDriver(c.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	// Set the ID
	resource.ResourceID = uuid.Must(uuid.NewRandom())

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
