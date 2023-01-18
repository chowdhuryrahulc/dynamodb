package product

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/adapter"
	"github.com/chowdhuryrahulc/dynamodb/internal/entities/product"
	"github.com/google/uuid"
)

// routes-->handlers-->controllers-->repository/adapters(contains database func, adapters means connect to db func)

type Controller struct{
	repository adapter.Interface	// comes from internal/repository/adapters/adpter.go
	// gives access to methords like Health, FindAll, FindOne, CreateOrUpdate, Delete
	// all these are database struct methods
	// with interfaces, other structs can use your methords. 
	// You just have to register your interface in other structs like done above
	// but first you need to add all the struct methords in the Interface interface{} before transporting it to other struct



}

type Interface interface{
	// Interfaces helps us create a lot of abstrction
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entities []product.Product, err error) 
	Create(entity *product.Product) (uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

type NewController(repository adapter.Interface) Interface{
	// this returns type Interface interface{} as writen above
	// this creates a new controller
	return &Controller{
		repository: repository
	}
}

// Below functions has 1:1 relationship with the routes and handlers

func (c *Controller) ListOne (id uuid.UUID) (entity product.Product, err error) {
	// listone needs uuid to find, and returns a product model
	entity.ID = id
	// entity.GetFilterId(): helps us get something from db directly
	// entity.TableName(): in which table does the value resides
	// FindOne: comes from adapters.adapters.go. It is a adapter interface function 
	// (means talks to database directly, controller-->adapter/repository)
	response, err:= c.repository.FindOne(entity.GetFilterId(), entity.TableName())  
	if err != nil {
		return entity, err
	}
	return product.ParseDynamoAttributeToStruct(response.Item) 
}

func (c *Controller) ListAll () (entities []product.Product, err error) {
	// listall gets all produt. So output is a list of all products
	entities = []product.Product{}				// multiple product (entity)
	var entity product.Product					// single product (entity)

	// we create a filtering variable that helps us to filter, and we set up filter variable to db, so we could filter based on a condition
	filter := expression.Name("name").NotEqual(expression.Value(""))	// setting up the filter
	// this means the product name should not be empty

	// this is the condition with which we need to filter (in dynamodb)
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return entities, err
	}

	//? filter --builds-> condition --> runs query in dynamodb

	//use repository to run the db level function (controllers-->repository)
	response, err := c.repository.FindAll(condition, entity.TableName()) // c is from the interface registered above, from repository folder
	// this function directly interacts with dynamodb	|| tablename: the table where we should run the query
	if err != nil {
		return entities, err
	}
	if response != nil {
		for _, value := range response.Items {
			// ParseDynamoAttributeToStruct: converts everything that is stored in dynamodb 
			// to struct that is understood by golang
			entity, err := product.ParseDynamoAttributeToStruct()
			if err != nil {
				return entities, err
			}
			entities = append(entities, entity)
		}
	}
	return entities, nil

}

func (c *Controller) Create (entity *product.Product) (uuid.UUID, error) {
	// gets the product model, creates the record, and returns the uuid of the product
	entity.CreatedAt = time.Now()	// we modify the createdAt value
	c.repository.CreateOrUpdate(entity.GetMap(), entity.TableName())	// from repository
	return entity.ID, err
}

func (c *Controller) Update (id uuid.UUID, entity *product.Product) error{
	// you get uuid and product model, and update in the position uuid with the new product model
	// do update func in the end, after completing all others


	
}

func (c *Controller) Remove (id uuid.UUID) error {
	// listone means you find that item in db
	entity, err := c.ListOne(id)
	if err != nil{
		return err
	} 
	_, err := c.repository.Delete(entity.GetFilterId(), entity.TableName())
	return err
}






