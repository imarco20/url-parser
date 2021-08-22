package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// logError prints the stack trace of an error
func (app *application) logError(r *http.Request, err error) {

	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.logger.Println(trace)
}

// errorResponse encodes the parameter message to the response writer
// and sets the response status code with the parameter value
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// serverErrorResponse is used in scenarios where the server encounters an error
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse sends a resource not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse is used in scenarios where the handles can't
// handle requests with some HTTP methods with a MethodNotAllowed status code
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// badRequestResponse is used in cases when the received request doesn't match
// the one that the handler expects
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
