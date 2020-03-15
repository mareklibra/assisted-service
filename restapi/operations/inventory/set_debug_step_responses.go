// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SetDebugStepOKCode is the HTTP code returned for type SetDebugStepOK
const SetDebugStepOKCode int = 200

/*SetDebugStepOK Registered

swagger:response setDebugStepOK
*/
type SetDebugStepOK struct {
}

// NewSetDebugStepOK creates SetDebugStepOK with default headers values
func NewSetDebugStepOK() *SetDebugStepOK {

	return &SetDebugStepOK{}
}

// WriteResponse to the client
func (o *SetDebugStepOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SetDebugStepInternalServerErrorCode is the HTTP code returned for type SetDebugStepInternalServerError
const SetDebugStepInternalServerErrorCode int = 500

/*SetDebugStepInternalServerError Internal server error

swagger:response setDebugStepInternalServerError
*/
type SetDebugStepInternalServerError struct {
}

// NewSetDebugStepInternalServerError creates SetDebugStepInternalServerError with default headers values
func NewSetDebugStepInternalServerError() *SetDebugStepInternalServerError {

	return &SetDebugStepInternalServerError{}
}

// WriteResponse to the client
func (o *SetDebugStepInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}