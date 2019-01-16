package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type App struct {
	Router     *mux.Router
	DriverInfo DriverInfo
}

func NewApp(driverInfo DriverInfo) *App {
	a := &App{
		DriverInfo: driverInfo,
	}
	a.initializeRoutes()
	return a
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/user", a.CreateUser).Methods("POST")
}

type User struct {
	EmailAddress string `json:"email_address"`
	FullName     string `json:"full_name"`
}

// CreateUser creates a user
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var resource User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resource); err != nil {
		requestUtils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	driver := GetDriver(a.DriverInfo)
	defer driver.Close()

	session := GetSession(driver)
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (User {name:$name, emailAddress:$emailAddress})",
			map[string]interface{}{"name": resource.FullName, "emailAddress": resource.EmailAddress})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	if err != nil {
		requestUtils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Created user"})
}

// ServeRest runs the server
func (a *App) ServeRest(addr, origin string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{origin})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "DELETE", "POST", "PUT", "OPTIONS"})
	log.Printf("Allowing origin: %s\n", origin)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, handlers.AllowCredentials(), headersOk, methodsOk)(a.Router)))
}
