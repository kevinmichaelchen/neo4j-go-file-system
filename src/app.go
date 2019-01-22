package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
)

type App struct {
	Router              *mux.Router
	DriverInfo          DriverInfo
	UserService         UserService
	OrganizationService OrganizationService
	MoveService         MoveService
	FileService         FileService
	FolderService       FolderService
}

func NewApp(driverInfo DriverInfo) *App {
	a := &App{
		DriverInfo:          driverInfo,
		UserService:         UserService{DriverInfo: driverInfo},
		OrganizationService: OrganizationService{DriverInfo: driverInfo},
		MoveService:         MoveService{DriverInfo: driverInfo},
		FileService:         FileService{DriverInfo: driverInfo},
		FolderService:       FolderService{DriverInfo: driverInfo},
	}
	a.initializeRoutes()
	return a
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/hello", HelloWorld).Methods(http.MethodGet)
	a.Router.HandleFunc("/user", a.UserService.CreateUser).Methods(http.MethodPost)
	a.Router.HandleFunc("/organization", a.OrganizationService.CreateOrganization).Methods(http.MethodPost)
	a.Router.HandleFunc("/move", a.MoveService.Move).Methods(http.MethodPost)
	a.Router.HandleFunc("/file/{id:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", a.FileService.GetFile).Methods(http.MethodGet)
	// TODO POST /org-membership w/ userID orgID
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
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