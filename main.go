package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
	"os"
)

func getDriver(driverInfo DriverInfo) neo4j.Driver {
	driver, err := neo4j.NewDriver(driverInfo.ConnectionUri, neo4j.BasicAuth(driverInfo.Username, driverInfo.Password, ""))
	if err != nil {
		panic(err)
	}
	return driver
}

func getSession(driver neo4j.Driver) neo4j.Session {
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		panic(err)
	}
	return session
}

func createObjects(session neo4j.Session) error {
	var (
		err    error
		result neo4j.Result
	)

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err = transaction.Run(
			`CREATE 
				(Kevin:User {name:'Kevin Chen'}),
				(IrisVR:Organization {name:'IrisVR'}),
				(IrisVRFolder1:Folder {name:'secret-projects'}),
				(File1:File {name:'File1.txt'}),
				(Kevin)-[:HAS_ORGANIZATION]->(IrisVR),
				(IrisVR)-[:CONTAINS_FOLDER]->(IrisVRFolder1),
				(IrisVRFolder1)-[:CONTAINS_FILE]->(File1)
`,
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	return err
}

func initializeObjects(driverInfo DriverInfo) {
	driver := getDriver(driverInfo)
	defer driver.Close()

	session := getSession(driver)
	defer session.Close()

	err := createObjects(session)
	if err != nil {
		panic(err)
	}
}

type DriverInfo struct {
	ConnectionUri string
	Username      string
	Password      string
}

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

	initializeObjects(DriverInfo{ConnectionUri: connectionString, Username: user, Password: pass})
}
