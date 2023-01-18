package health

import (
	"errors"
	"net/http"

	"github.com/chowdhuryrahulc/dynamodb/internal/handlers"
	"github.com/chowdhuryrahulc/dynamodb/internal/repository/adapter"
	HttpStatus "github.com/chowdhuryrahulc/dynamodb/utils/http"
)

type Handler struct {
	//todo What is meant by handlers.Interface. It is not a variable that is returned
	//todo adding Interface to a struct??
	handler.Interface				// comes from internal/handlers/interface.go
	Repository adapter.Interface	// comes from internal/repository/adapter
}

// All functions defined in interface.go defined below
func NewHandler(repository adapter.Interface) handler.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request){
	// for Health, we only need get. Bcoz we only need to check the health (but for product, we need get,post,put,delete etc)
	// thats why all post, put, delete, options in health(this file) just sends a http error back
	if !h.Repository.Health(){
		// errors comming from errors pkg
		// below line just sends a error response
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}
	HttpStatus.StatusOK(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request){
	HttpStatus.StatusMethodNotAllowed(w,r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request){
	HttpStatus.StatusMethodNotAllowed(w,r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request){
	HttpStatus.StatusMethodNotAllowed(w,r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request){
	HttpStatus.StatusNotContent(w,r)
}