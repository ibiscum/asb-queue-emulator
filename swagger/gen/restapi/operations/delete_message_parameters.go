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

// NewDeleteMessageParams creates a new DeleteMessageParams object
//
// There are no default values defined in the spec.
func NewDeleteMessageParams() DeleteMessageParams {

	return DeleteMessageParams{}
}

// DeleteMessageParams contains all the bound params for the delete message operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteMessage
type DeleteMessageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The token of the lock of the message to be deleted as returned by the Peek Message operation in BrokerProperties{LockToken}.
	  Required: true
	  In: path
	*/
	LockToken string
	/*The ID of the message to be deleted as returned in BrokerProperties{MessageId} by the Peek Message operation.
	  Required: true
	  In: path
	*/
	MessageID string
	/*the queue name
	  Required: true
	  In: path
	*/
	QueueName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteMessageParams() beforehand.
func (o *DeleteMessageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rLockToken, rhkLockToken, _ := route.Params.GetOK("lockToken")
	if err := o.bindLockToken(rLockToken, rhkLockToken, route.Formats); err != nil {
		res = append(res, err)
	}

	rMessageID, rhkMessageID, _ := route.Params.GetOK("messageId")
	if err := o.bindMessageID(rMessageID, rhkMessageID, route.Formats); err != nil {
		res = append(res, err)
	}

	rQueueName, rhkQueueName, _ := route.Params.GetOK("queueName")
	if err := o.bindQueueName(rQueueName, rhkQueueName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLockToken binds and validates parameter LockToken from path.
func (o *DeleteMessageParams) bindLockToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.LockToken = raw

	return nil
}

// bindMessageID binds and validates parameter MessageID from path.
func (o *DeleteMessageParams) bindMessageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.MessageID = raw

	return nil
}

// bindQueueName binds and validates parameter QueueName from path.
func (o *DeleteMessageParams) bindQueueName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.QueueName = raw

	return nil
}
