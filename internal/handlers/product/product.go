package product

import (
	"errors"
	"net/http"

	handler "github.com/chowdhuryrahulc/dynamodb/internal/handlers"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/adapter"
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
	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll() // this controller func gets all data from dynamodb
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.StatusOK(w, r, response)
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
	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})

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

func (h *Handler) getBodyAndValidate(w http.ResponseWriter, r *http.Request) () {
	// in post & put func you get a body from the data. That has to go inside dynamodb
	// this func validates if the operation is validate or not

}

func setDefaultValues(){

}