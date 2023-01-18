package product

import (
	"encoding/json"

	"github.com/chowdhuryrahulc/dynamodb/internal/entities"
)

type Product struct {
	entities.Base		// comes from entities/base.go, has id, createdAt, updatedAt (base struct and interface)
	Name string `json: "name"`
}

func InterfaceToModel(data interface{})(instance *Product, err error){
	// does marshaling & unmarshaing, converts json to struct (understandable by golang)
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}
	// UnMarshall: Converting JSON into Go objects
	return instance, json.Unmarshal(bytes, &instance) //❓
	//todo Why is unmarshal returning, when the return is error?
	//todo How is instance comming here?
}

func (p *Product) GetFilterId() map[string]interface{}{
	// used in controller findone function
	

}

func (p *Product) TableName()string{
	// returns name of table you want to send your data to, in dynamodb
	return "products"
}

func (p *Product) Bytes()([]byte, error){
	return json.Marshal(p)
}

func (p *Product) GetMap()map[string]interface{}{

}

func (p *Product) ParseDynamoAttributeToStruct()(){
	// dynamoAttribute-->struct

}



