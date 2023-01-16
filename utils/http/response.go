package http

import "net/http"

// most imp ile in utils

// defines the type of response the app will send
type response struct {
	Status int `json:"status"`			// like 500, 200 etc send by the server
	Result interface{} `json:"result"`
}

func newResponse()*response{

}

func (resp *response) bytes() []byte{}

func (resp *response) string() string{}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	//? most important function in this file
}

//* different response functions for different status codes below (great for production)
// if you dont say anything, backend will send only 500, which does not tell us the exact problem
//200
func StatusNoContent(){

}

//400
func StatusBadRequest(){}

//404
func StatusNotFound(){}

//405
func StatusMethodNotAllowed(){}

//409
func StatusConflict(){}

//500
func StatusInternalServerError(){}

