package rules

import (
	"io"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// we have 2 function that we will use through this interface
type Interface interface{
	ConvertIoReaderToStruct(data io.Reader, model interface{})(body interface{}, err error)
	GetMock() interface{}		// returns mock values
	Migrate(connection *dynamodb.DynamoDB) error
	Validate (model interface{}) error
}

 
