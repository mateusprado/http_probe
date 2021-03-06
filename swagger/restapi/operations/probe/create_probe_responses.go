// Code generated by go-swagger; DO NOT EDIT.

package probe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/mateusprado/http_probe/swagger/models"
)

// CreateProbeCreatedCode is the HTTP code returned for type CreateProbeCreated
const CreateProbeCreatedCode int = 201

/*CreateProbeCreated created

swagger:response createProbeCreated
*/
type CreateProbeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Probe `json:"body,omitempty"`
}

// NewCreateProbeCreated creates CreateProbeCreated with default headers values
func NewCreateProbeCreated() *CreateProbeCreated {

	return &CreateProbeCreated{}
}

// WithPayload adds the payload to the create probe created response
func (o *CreateProbeCreated) WithPayload(payload *models.Probe) *CreateProbeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create probe created response
func (o *CreateProbeCreated) SetPayload(payload *models.Probe) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProbeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
