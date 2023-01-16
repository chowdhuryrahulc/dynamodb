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

// below functions talk directly to dynamodb and get the result
// these functions will be called from the controllers

func (db *Database) Health() bool {
	// checks health of dynamodb
}

func (db *Database) FindAll {
	// checks health of dynamodb
}

func (db *Database) FindOne(condition map[string]interface{}, tableName) (response *dynamodb.GetItemOutput, err error) {
	// get 1 data from dynamodb
	conditionParsed, err := dynamodbattribute.MarshalMap(condition) // condition: means id=.., name=... type conditions. {} means no condition means get all the datsa
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName)	// dynamodb table from where you have to get the data
		Key: conditionParsed,
	}
	return db.Connection.GetItem(input)
}

func (db *Database) CreateOrUpdate {
	// checks health of dynamodb
}

func (db *Database) Delete {
	// checks health of dynamodb
}








