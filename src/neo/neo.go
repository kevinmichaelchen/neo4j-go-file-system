package neo

import (
	"fmt"
	"net/http"

	"github.com/kevinmichaelchen/neo4j-go-file-system/service"
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
				(IrisVR:Organization {name: 'IrisVR', resource_id: 'bdb77ea5-b0b2-4ecf-8605-4369b5b73577'}),
				(Kevin:User {full_name:'Kevin Chen', email_address: 'kevin.chen@irisvr.com', resource_id: 'fb59357e-a17e-45a4-9f2b-d967cf800c21'}),
				(Robin:User {full_name:'Robin Kim', email_address: 'robin@irisvr.com', resource_id: 'a08754ca-8a2c-4090-a761-54d8777f3ed4'}),
				(Graham:User {full_name:'Graham Hagger', email_address: 'graham@irisvr.com', resource_id: 'ce1d1610-4ad3-4fa1-9008-28aa86dbbeb9'}),
				(Ezra:User {full_name:'Ezra Smith', email_address: 'ezra@irisvr.com', resource_id: '1e018e77-1b66-4d95-8a1f-dba47ae190f4'}),
				(Shane:User {full_name:'Shane Scranton', email_address: 'shane@irisvr.com', resource_id: '3f1c0785-7b65-47fa-b9ab-243afe4fd8f5'}),
				(Nate:User {full_name:'Nate Beatty', email_address: 'nate@irisvr.com', resource_id: 'a5666488-d5fd-4ce9-ab47-702ece52b733'}),

				(CloudFolder:Folder {name:'cloud', resource_id: '0871b5af-4954-4d21-9e1f-3781e269374a'}),
				(CloudAuthFolder:Folder {name:'cloud-auth', resource_id: 'be94511e-c8b1-4e62-b37d-2e35704ea6c2'}),
				(CloudFileSystemFolder:Folder {name:'cloud-file-system', resource_id: '35be10e4-e114-4486-bc6e-d398f9793b42'}),
				(CloudFolder)-[:CONTAINS_FOLDER]->(CloudAuthFolder),
				(CloudFolder)-[:CONTAINS_FOLDER]->(CloudFileSystemFolder),
				(CloudAuthDoc:File {name:'cloud-auth.md', resource_id: '7a1ced19-5396-4c44-bc30-4953d59453d5'}),
				(CloudAuthFolder)-[:CONTAINS_FILE]->(CloudAuthDoc),

				(CloudLibrarySyncDoc:File {name:'cloud-library-syncing.md', resource_id: '9c73cde3-d8f9-4048-bfd9-00e0484fdb99'}),
				(CloudLibraryEventDataFormat:File {name:'cloud-library-event-data-format.md', resource_id: '05bb98d7-e088-4ffc-b685-04e4470b8a3a'}),
				(CloudFileSystemFolder)-[:CONTAINS_FILE]->(CloudLibrarySyncDoc),
				(CloudFileSystemFolder)-[:CONTAINS_FILE]->(CloudLibraryEventDataFormat),

				(AdminFolder:Folder {name:'admin', resource_id: '8257ce55-cbff-44ec-a50e-de721abc6aa2'}),
				(PayrollFile:File {name:'payroll.csv', resource_id: '48d85981-91db-427e-9c59-548c5a88acb7'}),
				(AdminFolder)-[:CONTAINS_FILE]->(PayrollFile),

				(DevOpsFolder:Folder {name:'devops', resource_id: 'cfb540b6-52b2-4469-8d8c-022bbe117e2f'}),
				(KubernetesDocFile:File {name:'kubernetes.md', resource_id: '4c2419c3-2546-42c9-bd66-4f47c53d281b'}),
				(DevOpsFolder)-[:CONTAINS_FILE]->(KubernetesDocFile),

				(Kevin)-[:HAS_ORGANIZATION]->(IrisVR),
				(Robin)-[:HAS_ORGANIZATION]->(IrisVR),
				(Graham)-[:HAS_ORGANIZATION]->(IrisVR),
				(Ezra)-[:HAS_ORGANIZATION]->(IrisVR),
				(Shane)-[:HAS_ORGANIZATION]->(IrisVR),
				(Nate)-[:HAS_ORGANIZATION]->(IrisVR),

				(IrisVR)-[:CONTAINS_FOLDER]->(CloudFolder),
				(IrisVR)-[:CONTAINS_FOLDER]->(AdminFolder),
				(IrisVR)-[:CONTAINS_FOLDER]->(DevOpsFolder),

				(Graham)-[:CAN_ACCESS]->(DevOpsFolder),
				(Shane)-[:CAN_ACCESS_EVERYTHING]->(IrisVR),
				(Nate)-[:CAN_ACCESS_EVERYTHING]->(IrisVR),
				(Kevin)-[:CAN_ACCESS]->(CloudFileSystemFolder),
				(Robin)-[:CAN_ACCESS]->(CloudAuthFolder),
				(Ezra)-[:CAN_ACCESS]->(CloudFolder)
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

func InitializeObjects(driverInfo DriverInfo) {
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

func RollbackIfError(tx neo4j.Transaction, originalError error) *service.Error {
	if originalError != nil {
		err := tx.Rollback()
		if err != nil {
			return service.NewError(http.StatusInternalServerError, fmt.Sprintf("Could not rollback transaction: %s", err.Error()), err)
		}
		return service.NewError(http.StatusInternalServerError, originalError.Error(), originalError)
	}
	return nil
}

type DriverInfo struct {
	ConnectionUri string
	Username      string
	Password      string
}
