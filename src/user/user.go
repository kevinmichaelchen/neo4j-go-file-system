package user

import (
	"encoding/json"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type User struct {
	ResourceID   uuid.UUID `json:"resourceID"`
	EmailAddress string    `json:"emailAddress"`
	FullName     string    `json:"fullName"`
}

func CreateUser(session neo4j.Session, user User) error {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			`CREATE (User {resource_id: $resource_id, email_address: $email_address, full_name: $full_name})`, userToMap(user))
	})
	if err != nil {
		return err
	}
	return nil
}

func userExists(session neo4j.Session, user User) (bool, error) {
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

func userToMap(user User) map[string]interface{} {
	return map[string]interface{}{
		"resource_id":   user.ResourceID.String(),
		"email_address": user.EmailAddress,
		"full_name":     user.FullName,
	}
}

type Service struct {
	DriverInfo neo.DriverInfo
}

// CreateUserRequestHandler creates a user
func (s *Service) CreateUserRequestHandler(w http.ResponseWriter, r *http.Request) {
	var resource User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := neo.GetDriver(s.DriverInfo)
	defer driver.Close()

	session := neo.GetSession(driver)
	defer session.Close()

	// Set the ID
	resource.ResourceID = uuid.Must(uuid.NewRandom())

	// TODO validate user resource
	exists, err := userExists(session, resource)
	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if exists {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "User already exists with that email")
		return
	}

	err = CreateUser(session, resource)

	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Created user"})
}
