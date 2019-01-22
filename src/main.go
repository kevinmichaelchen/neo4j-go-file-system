package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kevinmichaelchen/neo4j-go-file-system/neo"
)

func main() {
	var hostname, user, pass string
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
	connectionString := fmt.Sprintf("bolt://%s:7687", hostname)
	log.Printf("Connecting to: %s", connectionString)

	driverInfo := neo.DriverInfo{ConnectionUri: connectionString, Username: user, Password: pass}
	neo.InitializeObjects(driverInfo)

	a := NewApp(driverInfo)
	port := 8080
	a.ServeRest(fmt.Sprintf(":%d", port), "http://localhost:3000")
}
