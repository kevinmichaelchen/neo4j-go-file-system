package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
)

func main() {
	var hostname, user, pass, appPortString, grpcPortString string
	var ok bool
	if hostname, ok = os.LookupEnv("NEO_HOSTNAME"); !ok {
		log.Fatal("Failed to specify hostname")
	}
	if user, ok = os.LookupEnv("NEO_USERNAME"); !ok {
		log.Fatal("Failed to specify username")
	}
	if pass, ok = os.LookupEnv("NEO_PASSWORD"); !ok {
		log.Fatal("Failed to specify password")
	}
	if appPortString, ok = os.LookupEnv("APP_PORT"); !ok {
		log.Fatal("Failed to specify app port")
	}
	if grpcPortString, ok = os.LookupEnv("GRPC_PORT"); !ok {
		log.Fatal("Failed to specify gRPC port")
	}

	appPort, err := strconv.Atoi(appPortString)
	if err != nil {
		log.Fatalf("Failed to specify valid app port: %s", appPortString)
	}

	grpcPort, err := strconv.Atoi(grpcPortString)
	if err != nil {
		log.Fatalf("Failed to specify valid gRPC port: %s", grpcPortString)
	}

	connectionString := fmt.Sprintf("bolt://%s:7687", hostname)
	log.Printf("Connecting to: %s", connectionString)

	driverInfo := neo.DriverInfo{ConnectionUri: connectionString, Username: user, Password: pass}
	neo.InitializeObjects(driverInfo)

	a := NewApp(driverInfo, grpcPort)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.GrpcServer.Run()

	wg.Add(1)
	go a.ServeRest(fmt.Sprintf(":%d", appPort), "http://localhost:3000")

	wg.Wait()
}
