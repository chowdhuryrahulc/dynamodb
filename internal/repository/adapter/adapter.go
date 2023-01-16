package adapter

import "github.com/aws/aws-sdk-go/service/dynamodb"

type Database struct {
	// thius struct is used to tsalk to db
	connection *dynamodb.DynamoDB
	logMode bool
}

type Interface interface {

}

type NewAdapter() Interface {
	
}

//! below functions talk directly to dynamodb and get the result
//! these functions will be called from the controllers

func (db *Database) Health() bool {
	// checks health of dynamodb
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	// if tables are getting listed, means dynamodb is working fine
	return err == nil
}

func (db *Database) FindAll {
	// checks health of dynamodb
}

func (db *Database) FindOne(condition map[string]interface{}, tableName) (response *dynamodb.GetItemOutput, err error) {
	// get 1 data from dynamodb

	conditionParsed, err := dynamodbattribute.MarshalMap(condition) 
	// condition: means id=.., name=... type conditions. {} means no condition means get all the data
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName)	// dynamodb table from where you have to get the data
		Key: conditionParsed,
	}
	return db.Connection.GetItem(input)
}

func (db *Database) CreateOrUpdate (entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	// this same function is used for creating or updating
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}

	input:= &dynamodb.PutItemInput{
		Item: entityParsed,
		TableName: aws.String(tableName)
	}
	return db.Connection.PutItem(input)
}

func (db *Database) Delete (condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	// delete is always similar to findone, bcoz we find a item using a condition, and delete it
	&dynamodb.DeleteItemInput{
	conditionParsed, err := dynamodbattribute.MarshalMap(condition) 
	// condition: means id=.., name=... type conditions. {} means no condition means get all the data
	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		Key: conditionParsed,
		TableName: aws.String(tableName)	// dynamodb table from where you have to delete the data
	}
	return db.Connection.DeleteItem(input)
	}
}








