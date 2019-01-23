package neo

import (
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"
	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Service struct {
	DriverInfo neo.DriverInfo
}

func NewService(driverInfo neo.DriverInfo) *Service {
	return &Service{DriverInfo: driverInfo}
}

func (s *Service) CreateOrganization(resource organization.Organization) (*organization.Organization, *service.Error) {
	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	// TODO validate org resource

	exists, err := organizationExists(session, resource)
	if err != nil {
		return nil, service.Internal(err)
	}
	if exists {
		return nil, service.NewError(http.StatusBadRequest, "Org already exists with that name", nil)
	}

	err = createOrganization(session, resource)

	if err != nil {
		return nil, service.Internal(err)
	}

	return &resource, nil
}

func (s *Service) GetOrganization(resource organization.Organization) (*organization.Organization, *service.Error) {
	return nil, service.Unimplemented()
}

func (s *Service) UpdateOrganization(resource organization.Organization) (*organization.Organization, *service.Error) {
	return nil, service.Unimplemented()
}

func (s *Service) DeleteOrganization(resource organization.Organization) (*organization.Organization, *service.Error) {
	return nil, service.Unimplemented()
}

func createOrganization(session neo4j.Session, organization organization.Organization) error {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			`CREATE (Organization {resource_id: $resource_id, name: $name})`, orgToMap(organization))
	})
	if err != nil {
		return err
	}
	return nil
}

func organizationExists(session neo4j.Session, organization organization.Organization) (bool, error) {
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

func orgToMap(organization organization.Organization) map[string]interface{} {
	return map[string]interface{}{
		"resource_id": organization.ResourceID,
		"name":        organization.Name,
	}
}
