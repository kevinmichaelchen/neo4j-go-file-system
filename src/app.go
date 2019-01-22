package main

import (
	"log"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	"github.com/kevinmichaelchen/neo4j-go-file-system/move"
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
)

type App struct {
	Router              *mux.Router
	DriverInfo          neo.DriverInfo
	UserService         user.Service
	OrganizationService organization.Service
	MoveService         move.Service
	FileService         file.Service
	FolderService       folder.Service
}

func NewApp(driverInfo neo.DriverInfo) *App {
	a := &App{
		DriverInfo:          driverInfo,
		UserService:         user.Service{DriverInfo: driverInfo},
		OrganizationService: organization.Service{DriverInfo: driverInfo},
		MoveService:         move.Service{DriverInfo: driverInfo},
		FileService:         file.Service{DriverInfo: driverInfo},
		FolderService:       folder.Service{DriverInfo: driverInfo},
	}
	a.initializeRoutes()
	return a
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/hello", HelloWorldRequestHandler).Methods(http.MethodGet)
	a.Router.HandleFunc("/user", a.UserService.CreateUserRequestHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/organization", a.OrganizationService.CreateOrganizationRequestHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/move", a.MoveService.MoveRequestHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/file/{id:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", a.FileService.GetFileRequestHandler).Methods(http.MethodGet)
	// TODO POST /org-membership w/ userID orgID
}

func HelloWorldRequestHandler(w http.ResponseWriter, r *http.Request) {
	requestUtils.RespondWithJSON(w, http.StatusOK, map[string]string{"msg": "Hello world"})
}

// ServeRest runs the server
func (a *App) ServeRest(addr, origin string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{origin})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "DELETE", "POST", "PUT", "OPTIONS"})
	log.Printf("Allowing origin: %s\n", origin)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, handlers.AllowCredentials(), headersOk, methodsOk)(a.Router)))
}
