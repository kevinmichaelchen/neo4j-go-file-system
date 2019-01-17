package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	EmailAddress string    `json:"email_address"`
	FullName     string    `json:"full_name"`
}

func CreateUser(session neo4j.Session, user User) error {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		return transaction.Run(
			`CREATE (User {id: $id, email_address: $email_address, full_name: $full_name})`, userToMap(user))
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
		log.Println("RESULT =", e)
		return e != "", nil
	}
	return false, nil
}

func userToMap(user User) map[string]interface{} {
	return map[string]interface{}{
		"id":            user.ID.String(),
		"email_address": user.EmailAddress,
		"full_name":     user.FullName,
	}
}

type UserService struct {
	DriverInfo DriverInfo
}

// CreateUser creates a user
func (s *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var resource User
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
