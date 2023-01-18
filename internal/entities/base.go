package entities 
// base.go has id, createdAt, updatedAt etc. These are base things 
// product.go has the remaining product model 
// so base.go is a part of product.go
// most files follow pattern: struct->interface->methods	(only for this project)

import (
	"github.com/google/uuid"
	"time"						// used for creating createdAt & updatedAt
)

type Interface interface{
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	TableName() string
	GetMap() map[string]interface{}
	GetFilterId() map[string]interface{}
}

type Base struct {
	ID				uuid.UUID	`json:"_id"`
	CreatedAt		time.Time	`json:"createdAt"`
	UpdatedAt		time.Time	`json:"updatedAt"`
}

// below are the base struct methods to set id, createdat, updatedat, also defined in interfaces above

func (b *Base) GenerateId(){
	b.ID = uuid.New() // this methord creates a new id
}

func (b *Base) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}

func GetTimeFormat()string{
	// sets the format in which you want your time pkg to return time
	return "2010-01-02T15:04:05-0700" // this is regular timestamp, get from google (constant accross the world)
}

// func (b *Base) TableName() string {

// }

// func (b *Base) GetMap() map[string]interface{

// }

// func (b *Base) GetFilterId() map[string]interface{}
