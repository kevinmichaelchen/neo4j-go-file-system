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
	var hostname, user, pass, appPortString, internalGrpcPortString, externalGrpcPortString string
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
	if internalGrpcPortString, ok = os.LookupEnv("INTERNAL_GRPC_PORT"); !ok {
		log.Fatal("Failed to specify gRPC port")
	}
	if externalGrpcPortString, ok = os.LookupEnv("EXTERNAL_GRPC_PORT"); !ok {
		log.Fatal("Failed to specify gRPC port")
	}

	appPort, err := strconv.Atoi(appPortString)
	if err != nil {
		log.Fatalf("Failed to specify valid app port: %s", appPortString)
	}

	internalGrpcPort, err := strconv.Atoi(internalGrpcPortString)
	if err != nil {
		log.Fatalf("Failed to specify valid gRPC port: %s", internalGrpcPortString)
	}

	externalGrpcPort, err := strconv.Atoi(externalGrpcPortString)
	if err != nil {
		log.Fatalf("Failed to specify valid gRPC port: %s", externalGrpcPortString)
	}

	connectionString := fmt.Sprintf("bolt://%s:7687", hostname)
	log.Printf("Connecting to: %s", connectionString)

	driverInfo := neo.DriverInfo{ConnectionUri: connectionString, Username: user, Password: pass}
	neo.InitializeObjects(driverInfo)

	a := NewApp(driverInfo, internalGrpcPort, externalGrpcPort)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.InternalGrpcServer.Run()

	wg.Add(1)
	go a.ExternalGrpcServer.Run()

	wg.Add(1)
	go a.ServeRest(fmt.Sprintf(":%d", appPort), "http://localhost:3000")

	wg.Wait()
}
