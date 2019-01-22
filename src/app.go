package main

import (
	"log"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/grpc"

	"github.com/kevinmichaelchen/neo4j-go-file-system/file"
	fileNeo "github.com/kevinmichaelchen/neo4j-go-file-system/file/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/folder"
	folderNeo "github.com/kevinmichaelchen/neo4j-go-file-system/folder/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/move"
	moveNeo "github.com/kevinmichaelchen/neo4j-go-file-system/move/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/organization"
	orgNeo "github.com/kevinmichaelchen/neo4j-go-file-system/organization/neo"
	"github.com/kevinmichaelchen/neo4j-go-file-system/user"
	userNeo "github.com/kevinmichaelchen/neo4j-go-file-system/user/neo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	requestUtils "github.com/kevinmichaelchen/my-go-utils/request"
	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
)

type App struct {
	Router                 *mux.Router
	GrpcServer             grpc.Server
	DriverInfo             neo.DriverInfo
	UserController         user.Controller
	OrganizationController organization.Controller
	MoveController         move.Controller
	FileController         file.Controller
	FolderController       folder.Controller
}

func NewApp(driverInfo neo.DriverInfo, grpcPort int) *App {
	userService := userNeo.NewService(driverInfo)
	organizationService := orgNeo.NewService(driverInfo)
	moveService := moveNeo.NewService(driverInfo)
	fileService := fileNeo.NewService(driverInfo)
	folderService := folderNeo.NewService(driverInfo)

	a := &App{
		GrpcServer: grpc.Server{
			Port:                grpcPort,
			UserService:         userService,
			OrganizationService: organizationService,
			MoveService:         moveService,
			FileService:         fileService,
			FolderService:       folderService,
		},
		DriverInfo:             driverInfo,
		UserController:         user.Controller{Service: userService},
		OrganizationController: organization.Controller{Service: organizationService},
		MoveController:         move.Controller{Service: moveService},
		FileController:         file.Controller{Service: fileService},
		FolderController:       folder.Controller{Service: folderService},
	}
	a.initializeRoutes()
	a.GrpcServer.Run()
	return a
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/hello", HelloWorldRequestHandler).Methods(http.MethodGet)
	a.Router.HandleFunc("/user", a.UserController.CreateUserRequestHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/organization", a.OrganizationController.CreateOrganizationRequestHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/move", a.MoveController.MoveRequestHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/file/{id:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", a.FileController.GetFileRequestHandler).Methods(http.MethodGet)
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
