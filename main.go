package main

import (
	"fmt"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func GetDriver(driverInfo DriverInfo) neo4j.Driver {
	driver, err := neo4j.NewDriver(driverInfo.ConnectionUri, neo4j.BasicAuth(driverInfo.Username, driverInfo.Password, ""))
	if err != nil {
		panic(err)
	}
	return driver
}

func GetSession(driver neo4j.Driver) neo4j.Session {
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
				(Kevin:User {name:'Kevin Chen', emailAddress: 'kevin.chen@irisvr.com'}),
				(Robin:User {name:'Robin Kim', emailAddress: 'robin@irisvr.com'}),
				(Graham:User {name:'Graham Hagger', emailAddress: 'graham@irisvr.com'}),
				(Ezra:User {name:'Ezra Smith', emailAddress: 'ezra@irisvr.com'}),
				(Shane:User {name:'Shane Scranton', emailAddress: 'shane@irisvr.com'}),
				(Nate:User {name:'Nate Beatty', emailAddress: 'nate@irisvr.com'}),
				(IrisVR:Organization {name:'IrisVR'}),
				(SecretsFolder:Folder {name:'secret-projects'}),
				(AdminFolder:Folder {name:'admin'}),
				(DevOpsFolder:Folder {name:'devops'}),
				(NextGenVRFile:File {name:'next-gen-vr.txt'}),
				(PayrollFile:File {name:'payroll.csv'}),
				(KubernetesDocFile:File {name:'kubernetes.md'}),
				(Kevin)-[:HAS_ORGANIZATION]->(IrisVR),
				(Robin)-[:HAS_ORGANIZATION]->(IrisVR),
				(Graham)-[:HAS_ORGANIZATION]->(IrisVR),
				(Ezra)-[:HAS_ORGANIZATION]->(IrisVR),
				(Shane)-[:HAS_ORGANIZATION]->(IrisVR),
				(Nate)-[:HAS_ORGANIZATION]->(IrisVR),
				(IrisVR)-[:CONTAINS_FOLDER]->(SecretsFolder),
				(IrisVR)-[:CONTAINS_FOLDER]->(AdminFolder),
				(IrisVR)-[:CONTAINS_FOLDER]->(DevOpsFolder),
				(SecretsFolder)-[:CONTAINS_FILE]->(NextGenVRFile),
				(AdminFolder)-[:CONTAINS_FILE]->(PayrollFile),
				(DevOpsFolder)-[:CONTAINS_FILE]->(KubernetesDocFile)
`,
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})

	return err
}

func deleteAll(session neo4j.Session) {
	_, err := session.Run("MATCH (n) DETACH DELETE n", map[string]interface{}{})
	if err != nil {
		panic(err)
	}
}

func initializeObjects(driverInfo DriverInfo) {
	driver := GetDriver(driverInfo)
	defer driver.Close()

	session := GetSession(driver)
	defer session.Close()

	//deleteAll(session)

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

	driverInfo := DriverInfo{ConnectionUri: connectionString, Username: user, Password: pass}
	initializeObjects(driverInfo)

	a := NewApp(driverInfo)
	port := 8080
	a.ServeRest(fmt.Sprintf(":%d", port), "http://localhost:3000")
}
