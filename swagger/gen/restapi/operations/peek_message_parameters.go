// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewPeekMessageParams creates a new PeekMessageParams object
//
// There are no default values defined in the spec.
func NewPeekMessageParams() PeekMessageParams {

	return PeekMessageParams{}
}

// PeekMessageParams contains all the bound params for the peek message operation
// typically these are obtained from a http.Request
//
// swagger:parameters peekMessage
type PeekMessageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the queue name
	  Required: true
	  In: path
	*/
	QueueName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPeekMessageParams() beforehand.
func (o *PeekMessageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rQueueName, rhkQueueName, _ := route.Params.GetOK("queueName")
	if err := o.bindQueueName(rQueueName, rhkQueueName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindQueueName binds and validates parameter QueueName from path.
func (o *PeekMessageParams) bindQueueName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.QueueName = raw

	return nil
}
