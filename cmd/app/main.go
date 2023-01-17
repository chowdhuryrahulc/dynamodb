package main

/*
PROJECT STRUCTURE:
migrate func	: migrate dynamodb tables ❓					NECESSARY FOR BERLINGER
check tbles func: checks if dynamodb tables exist or not	 NECESSARY FOR BERLINGER

routes-->handlers-->controllers-->repository/adapters(contains database func, adapters means connect to db func)

helpers/rules (these func can be used by any file in the project)
	1) ConvertIoReaderToStruct: converts the data that is comming to struct to be used by golang
	2) Migrate: migrate func
	3) Create Tables: creates a new tble
	4) Validate: validate db level

entities: (every product will have its own id, createdat, updateat, tablename and name, something like a model)
	1) Base		: id, createdat, updateat, tablename
	2) Product	: base, name

json-->struct-->dynamo attribute-------------
		^---parse dynao attibute to struct<--
	json gets changed to struct(bcoz golang understands struct, not json) which
	then gets converted into dynamo attribute(bcoz dynamodb interacts with its own attributes)
	json is send by postman, or curl from terminal
	multiple functions will be used to change btw json, struct, dynamo attribute
	eg: parse dynao attibute to struct func: changes from dynamo attribute to struct

handlers --> interface(controller level) --> ListOne, Update, ListAll, Remove, Create (these are functions)
Controller level-->interface(database level, repository/adapters)-->FindOne, Delete, FindAll, Create/Update
	routes give control to handlers,
	When handlers talk to controllers, they talk using interfacs
	When controllers talk to databases, they talk using interfacs
	controller func and database func have 1:1 relationship.
	Means ListOne:FindOne, Update:Create/Update, Remove:Delete etc. (ListOne calls FindOne, etc type relationship)
	(1 controller func uses 1 db func)

Post/Put request format:
	Process1: (Best formt, Not implemented in this project)
		?body(json format)-->unmarshal()-->validate-->marshal-->store in db(dynamodb)
		golang does not understand json
	Process2: (using Interfaces, implemented in this project)
		body-->struct-->interface to model-->validate
! Dont implement 2nd process, as done here. Use 1st, as done in other projects
! Also use project folder structure used in other projects (cmd and pkg) instead of here (flask/django format)
		todo done in complete serverless ... video (dynamodb used)


***************************************************************************************************

In this project (chi router)
1) using interfaces
2) use logger
3) handling CORS
4) Health Check
5) Recover middleware in chi
6) Ozzo validation
pkg & cmd (used in mysql project) was a better project structure


Repository: has code about talking to db
	1) adapter-->adapter.go-->
	2) instance-->instance.go--> started dynamodb session

	create foundation of project by creating routes, enviournment, config, response,logger etc


*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/chowdhuryrahulc/dynamodb/config"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/adapter"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/instance"
	"github.com/chowdhuryrahulc/dynamodb/internal/routes"
	"github.com/chowdhuryrahulc/dynamodb/internal/rules"
	"github.com/chowdhuryrahulc/dynamodb/internal/rules/product"
	"github.com/chowdhuryrahulc/dynamodb/utils/logger"
)


func main(){
	configs := config.GetConfig()			// from config file
	connection := instance.GetConnection()	// sets up a connection/session with dynamodb (from internal/repository/instance)
	repository := adapter.NewAdapter(connection) // returns the Database struct. You can access entire database using this (from internam/repository/adapter) 

	logger.INFO("waiting for the service to start.....", nil)
	errors := Migrate(connection)			// for database migration (can be skipped)(implemented below)
	if len(errors)>0{						// logging database migration errors
		for _, err := range errors{
			logger.PANIC("Error on migration:....", err)
		}
	}

	logger.PANIC("", checkTables(connection)) // logging check table function errors 
	
	port := fmt.Sprintf("%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository) //todo What does this do? (from routers folder)
	logger.INFO("service is running on port", port)

	error:= http.ListenAndServe(port, router) // creates a server
	log.Fatal(error)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	//todo for database migration (research more about it)
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface){
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	// connection = instance.getconnection which connects to dynamodb session
	// ListTables lists all the tables you have in your dynamodb
	// if the number of tables is 0, then it shows no dynamodb tables found. 
	// But if the number of tables in dynamodb is more than 1, then it says tables found
	// BERLINGER: I think it will be better to check name of the dynamodb tables. 
	// And when we change or update the dynamodb table to a new table, it should auto-detect an start reading from the new table
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames)== 0{
			logger.INFO("Tables not found:", nil)
		}
		for _, tableName := range response.TableNames {
			logger.INFO("Table found:", tableName)
		}
	}
	return err
}