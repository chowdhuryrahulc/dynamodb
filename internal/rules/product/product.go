package product

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/chowdhuryrahulc/dynamodb/internal/entities"
	"github.com/chowdhuryrahulc/dynamodb/internal/entities/product"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

// In this file we convert one datatype to another (json->struct->dynamoattribute etc)
type Rules struct{}

func NewRules() *Rules {
	return &Rules{} // returns empty rules
}

// We will interact with the functions present in rules/interface.go from this file. Thats why we added them all in interface (in interfce.go)
func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	// in this func, we get some data and decode it into a model

	if data == nil { // return error if no data is recieved
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model) // we decode the incoming data into a model
}

func (r *Rules) GetMock() interface{} { // returns mock values
	// (returns product interface with some default/garbage values)

	return product.Product{ // comes from entities/product/product.go
		Base: entities.Base{ // comes from entities/base.go
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.New().String(),
	}
}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	return r.CreateTable(connection)
}

func (r *Rules) Validate(model interface{}) error {
	// change from interface --> model
	productModel, err := product.InterfaceToModel(model)
	if err != nil {
		return err
	}

	// below line just validates id & name (comes from....not done yet)
	return Validation.ValidateStruct(productModel,
		Validation.Field(&productModel.ID, Validation.Required, is.UUIDv4),
		Validation.Field(&productModel.Name, Validation.Required, Validation.Length(3, 50)),
	)

}

func (r *Rules) CreateTable(connection *dynamodb.DynamoDB) error {
	// this func is not in interfce.go, bcoz migrate func is calling this func, not the user
	// The whole code below is taken from dynamodb docomentation directly
	// so copy the below code in every dynamodb project
	// you have to use this func every time you connect to dynamodb

	//**********************************************************************************************************************************
	//									PERSONAL NOTES
	//! Where is this func used?
	//Sol: Migrate, in this file And migrate is used in main.go and internal/handlers/product/product.go
	// In main.go it is used in Migrate function, which is then used in main.go as [errors := Migrate(connection)]
	// this is used as part of database migration, still todo
	// todo Learn more about database migration
	// in handlers/products/products.go it is giving error in RulesProduct.NewRules().
	// Means probably bcoz of interface this error is present?? In interface Migrate methord is mentioned. So that gives error
	// And so..❓
	// Even in main, RulesProduct.NewRules() shows error: Sol-> see interface repo. It create a new instance in different file
	//todo And where does the connection go??
	//!Why is that if we delete any of the methord from rules, RulesProduct.NewRules() shows error? Where is that comming from??
	// Sol: see interface repository. It basically creates a new instance for that file to use
	//todo Learn this project & golang basics
	//**********************************************************************************************************************************

	//todo And where to put it in BERLINGER
	table := &product.Product{}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(table.TableName()),
	}
	response, err := connection.CreateTable(input)
	if err != nil && strings.Contains(err.Error(), "Table already exists") {
		//* This should be where the func should always return nil, bcoz table will always exist
		return nil
	}

	if response != nil && strings.Contains(response.GoString(), "TableStatus:\"CREATING\"") {
		time.Sleep(3 * time.Second)
		err = r.CreateTable(connection)
		if err != nil {
			return err
		}
	}
	return nil
}
