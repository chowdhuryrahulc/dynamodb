package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/chowdhuryrahulc/dynamodb/internal/controllers"//If error,add /products
	EntityProduct "github.com/chowdhuryrahulc/dynamodb/internal/entities/product"
	handler "github.com/chowdhuryrahulc/dynamodb/internal/handlers"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/adapter"
	Rules "github.com/chowdhuryrahulc/dynamodb/internal/rules"
	RulesProduct "github.com/chowdhuryrahulc/dynamodb/internal/rules/product"
	HttpStatus "github.com/chowdhuryrahulc/dynamodb/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

//! All below are handler functions. routes-->handlers-->controller-->Repository/Adapters
//! So all handler function calls something from controller (after doing some validation)
// Controllers are the MOST IMP
// Repository/Adapters--> does all the CRUD operations on Dynamodb(or other db)
// So we know where to look for errors, validations, etc

// get api: checks health of entire project (in handlers/health/health.go)
// get, getOne, getAll, post, put etc: apis related to product (in this file, handlers/product/product.go)

// creating handlers for products
type Handler struct {

	//******************************************************************************************************************
	// 								PERSONAL NOTES
	// 3 interfces, handler, product and rules added in this struct
	// handler interface				: get, post, put, delete, options
	// controller or product interface	: listone, listall, create, update, remove
	// rules interface					: convertIoReaderToStruct, getmock, migrate, validate
	//******************************************************************************************************************


	handler.Interface // using interfaces (mentioned in internal/handlers/product.go)
	Controller        product.Interface
	Rules             Rules.Interface
}

func NewHandler(repository adapter.Interface) handler.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// this func will determine if the control goes to getOne or getAll function
	// if ID gets passed in params, then we will call getOne function. If it doesnt, call getAll

	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

// this function just gets 1 response from the dynamodb based on id and sends back to user
func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	// routes->handlers(after doing some validation)->controllers

	// uuid comes from uuid pkg
	ID, err := uuid.Parse(chi.URLParam(r, "ID")) // we get id from params
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	// controller func, which gets 1 record from dynamodb
	response, err := h.Controller.ListOne(ID)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOk(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll() // this controller func gets all data from dynamodb
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOk(w, r, response)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	// we validate the request body to get the product body
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}
	// After you create something, you get a ID back (see controllers folder/create func)
	ID, err := h.Controller.Create(productBody)

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOk(w, r, map[string]interface{}{"id": ID.String()})

}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	// put methord is the most difficult than other methords
	ID, err := uuid.Parse(chi.URLParam(r, "ID")) // we get id from params
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	// validation step
	productBody, err := h.getBodyAndValidate(r, ID)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	// in this controller function, we find the data from dynamodb database using id and update it (in Controller update func)
	if err := h.Controller.Update(ID, productBody); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusNoContent(w,r)

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// we get id and then delete the record based on id
	ID, err := uuid.Parse(chi.URLParam(r, "ID")) // we get id from params
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}
	if err := h.Controller.Remove(ID); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusNoContent(w,r)


}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w,r)
}

// This function is used in Post and Put methods to work with body we get in the request
//* Put/Post method used: body-->struct-->interface to model-->validate (done in this func) (not for Berlinger, Process1 was better for Berlinger)
// 		this func also sets default values of CreatedAt/UpdatedAt
func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityProduct.Product, error) { // EntityProduct.Product comes from entities folder
	// in post & put func you get a body from the data. That has to go inside dynamodb
	// this func validates if the operation is validate or not
	productBody := &EntityProduct.Product{}	//todo What does EntityProduct.Product{} do?

	// converting body(json format) to struct/interface type (comes from rules folder)
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	// changing struct/interface into model
	productParsed, err := EntityProduct.InterfaceToModel(body)
	if err != nil{
		return &EntityProduct.Product{}, errors.New("error in converting body to model")
	}

	setDefaultValues(productParsed, ID)	// you want to update CreatedAt, UpdatedAt values of the api

	// we return the validated result
	return productParsed, h.Rules.Validate(productParsed)

}

func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID){
	// update CreatedAt, UpdatedAt values of the api
	product.UpdatedAt = time.Now()
	if ID == uuid.Nil{					// uuid.Nil is send only by POST request. For put, we dont have to set created at. Only updatedAt
		product.ID = uuid.New()			// New uuid is given to the new product to be send to db
		product.CreatedAt = time.Now() 	// creation time of this new prouct is now
	} else {
		product.ID = ID					// only for PUT
	}
}