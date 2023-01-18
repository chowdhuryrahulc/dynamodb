package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// most imp ile in utils

type response struct {
	// defines the type of response the app will send to front-end
	Status int         `json:"status"` // like 500, 200 etc send by the server
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *response {
	// whenever we want to send a response to the front-end, we call this function
	// this func is used in all status codes
	return &response{
		Status: status,
		Result: data,
	}
}

func (resp *response) bytes() []byte {
	// this method just marshels the data
	data, _ := json.Marshal(resp)
	return data
}

func (resp *response) string() string {
	// this converts data to string
	return string(resp.bytes())
}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	//? most important function in this file
	// here we send response to front-end
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")	// this handles CORS errors
	w.WriteHeader(resp.Status)
	_,_ = w.Write(resp.bytes())
	log.Println(resp.string())

}

//* different response functions for different status codes below (great for production)
// if you dont say anything, backend will send only 500, which does not tell us the exact problem
// below are all status errors
//200
func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}) {
	// data & http.StatusOK gets converted to response struct in newResponse
	// and that response gets send to sendResponse which sends the response
	newResponse(data, http.StatusOK).sendResponse(w,r)
}

//204
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	//NoContent error means no data interface{} will be passed
	newResponse(nil, http.StatusNoContent).sendResponse(w,r)
}

//400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data:= map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w,r)
}

//404
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data:= map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w,r)
}

//405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w,r)
}

//409
func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data:= map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusOK).sendResponse(w,r)
}

//500
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data:= map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusOK).sendResponse(w,r)
}
