package main

/*
PROJECT STRUCTURE:
migrate func	: migrate dynamodb tables ❓
check tbles func: checks if dynamodb tables exist or not

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





func main(){
	
}