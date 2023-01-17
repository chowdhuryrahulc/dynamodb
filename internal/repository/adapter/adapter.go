package adapter	//! This file has functions that talks directly to dynamodb

import (
	// take care of these specific imports
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Database struct {
	// thus struct is used to talk to db
	connection *dynamodb.DynamoDB
	logMode bool
}

type Interface interface {
	// this is a list of all the functions we have below, db struct methords
	Health() bool
	FindAll(condition expression.Expression, tableName string)(response *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tableName string)(response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string)(response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string)(response *dynamodb.DeleteItemOutput, err error)
}

func NewAdapter(con *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: con,
		logMode: false,
	}
}

//! below functions talk directly to dynamodb and get the result
//! these functions will be called from the controllers

func (db *Database) Health() bool {
	// checks health of dynamodb
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	// if tables are getting listed, means dynamodb is working fine
	return err == nil
}

func (db *Database) FindAll (condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	// expression.Expression is a part of dynamodb package
	input := dynamodb.ScanInput{
		ExpressionAttributeNames: condition.Names(),	// dynamodb query: conditions used to filter and fetch data
		ExpressionAttributeValues: condition.Values(),
		FilterExpression: condition.Filter(),
		ProjectionExpression: condition.Projection(),
		TableName: aws.String(tableName),
	}
}


func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	// get 1 data from dynamodb

	conditionParsed, err := dynamodbattribute.MarshalMap(condition) 
	// condition: means id=.., name=... type conditions. {} means no condition means get all the data
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),	// dynamodb table from where you have to get the data
		Key: conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate (entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	// this same function is used for creating or updating
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}

	input:= &dynamodb.PutItemInput{
		Item: entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) Delete (condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	// delete is always similar to findone, bcoz we find a item using a condition, and delete it
		conditionParsed, err := dynamodbattribute.MarshalMap(condition)
		// condition: means id=.., name=... type conditions. {} means no condition means get all the data
		if err != nil {
			return nil, err
		}

		input := &dynamodb.DeleteItemInput{
			Key: conditionParsed,
			TableName: aws.String(tableName),	// dynamodb table from where you have to delete the data
		}
		return db.connection.DeleteItem(input)
}