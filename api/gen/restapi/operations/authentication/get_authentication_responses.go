// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAuthenticationOKCode is the HTTP code returned for type GetAuthenticationOK
const GetAuthenticationOKCode int = 200

/*GetAuthenticationOK authentication successful

swagger:response getAuthenticationOK
*/
type GetAuthenticationOK struct {

	/*
	  In: Body
	*/
	Payload *GetAuthenticationOKBody `json:"body,omitempty"`
}

// NewGetAuthenticationOK creates GetAuthenticationOK with default headers values
func NewGetAuthenticationOK() *GetAuthenticationOK {

	return &GetAuthenticationOK{}
}

// WithPayload adds the payload to the get authentication o k response
func (o *GetAuthenticationOK) WithPayload(payload *GetAuthenticationOKBody) *GetAuthenticationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get authentication o k response
func (o *GetAuthenticationOK) SetPayload(payload *GetAuthenticationOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthenticationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
