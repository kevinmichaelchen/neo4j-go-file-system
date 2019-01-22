package neo

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type NeoService struct {
	DriverInfo neo.DriverInfo
}

func NewNeoService(driverInfo neo.DriverInfo) *NeoService {
	return &NeoService{DriverInfo: driverInfo}
}

func (s *NeoService) CreateUser(resource user.User) (*user.User, *user.ServiceError) {
	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	// Set the ID
	resource.ResourceID = uuid.Must(uuid.NewRandom())

	// TODO validate user resource
	exists, err := userExists(session, resource)
	if err != nil {
		return nil, &user.ServiceError{
			HttpCode:     http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			Error:        err,
		}
	}
	if exists {
		return nil, &user.ServiceError{
			HttpCode:     http.StatusBadRequest,
			ErrorMessage: "User already exists with that email",
			Error:        nil,
		}
	}

	err = createUser(session, resource)

	if err != nil {
		return nil, &user.ServiceError{
			HttpCode:     http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			Error:        err,
		}
	}

	return &resource, nil
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
		"resource_id":   user.ResourceID.String(),
		"email_address": user.EmailAddress,
		"full_name":     user.FullName,
	}
}
