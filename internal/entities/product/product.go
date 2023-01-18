package product

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/chowdhuryrahulc/dynamodb/internal/entities"
	"github.com/google/uuid"
)

type Product struct {
	entities.Base		// comes from entities/base.go, has id, createdAt, updatedAt (base struct and interface)
	Name string `json:"name"`
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
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Product) TableName()string{
	// returns name of table you want to send your data to, in dynamodb
	return "products"
}

func (p *Product) Bytes()([]byte, error){
	return json.Marshal(p)
}

func (p *Product) GetMap()map[string]interface{}{
	return map[string]interface{}{
		"_id":       p.ID.String(),
		"name":      p.Name,
		"createdAt": p.CreatedAt.Format(entities.GetTimeFormat()),
		"updatedAt": p.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}

// dynamoAttribute-->struct
func ParseDynamoAtributeToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return p, errors.New("Item not found")
	}
	for key, value := range response {
		if key == "_id" {
			p.ID, err = uuid.Parse(*value.S)
			if p.ID == uuid.Nil {
				err = errors.New("Item not found")
			}
		}
		if key == "name" {
			p.Name = *value.S
		}
		if key == "createdAt" {
			p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "updatedAt" {
			p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}