package neo

import (
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/service"

	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Service struct {
	DriverInfo neo.DriverInfo
}

func NewService(driverInfo neo.DriverInfo) *Service {
	return &Service{DriverInfo: driverInfo}
}

func (s *Service) CreateUser(resource user.User) (*user.User, *service.Error) {
	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	// TODO validate user resource
	exists, err := userExists(session, resource)
	if err != nil {
		return nil, service.Internal(err)
	}
	if exists {
		return nil, service.AlreadyExists("User already exists with that email")
	}

	err = createUser(session, resource)

	if err != nil {
		return nil, service.Internal(err)
	}

	return &resource, nil
}

func (s *Service) GetUser(resource user.User) (*user.User, *service.Error) {
	return nil, service.Unimplemented()
}

func (s *Service) UpdateUser(resource user.User) (*user.User, *service.Error) {
	return nil, service.Unimplemented()
}

func (s *Service) DeleteUser(resource user.User) (*user.User, *service.Error) {
	return nil, service.Unimplemented()
}

func createUser(session neo4j.Session, user user.User) error {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			`CREATE (User {resource_id: $resource_id, email_address: $email_address, full_name: $full_name})`, userToMap(user))
	})
	if err != nil {
		return err
	}
	return nil
}

func userExists(session neo4j.Session, user user.User) (bool, error) {
	res, err := session.Run(`MATCH (u:User {email_address: $email_address}) RETURN u.email_address`, map[string]interface{}{"email_address": user.EmailAddress})
	if err != nil {
		return false, err
	}
	if res.Next() {
		e := res.Record().GetByIndex(0).(string)
		return e != "", nil
	}
	return false, nil
}

func userToMap(user user.User) map[string]interface{} {
	return map[string]interface{}{
		"resource_id":   user.ResourceID,
		"email_address": user.EmailAddress,
		"full_name":     user.FullName,
	}
}
