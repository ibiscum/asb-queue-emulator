// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SendMessageCreatedCode is the HTTP code returned for type SendMessageCreated
const SendMessageCreatedCode int = 201

/*SendMessageCreated Message successfully sent to queue or topic.

swagger:response sendMessageCreated
*/
type SendMessageCreated struct {
}

// NewSendMessageCreated creates SendMessageCreated with default headers values
func NewSendMessageCreated() *SendMessageCreated {

	return &SendMessageCreated{}
}

// WriteResponse to the client
func (o *SendMessageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// SendMessageBadRequestCode is the HTTP code returned for type SendMessageBadRequest
const SendMessageBadRequestCode int = 400

/*SendMessageBadRequest Bad request.

swagger:response sendMessageBadRequest
*/
type SendMessageBadRequest struct {
}

// NewSendMessageBadRequest creates SendMessageBadRequest with default headers values
func NewSendMessageBadRequest() *SendMessageBadRequest {

	return &SendMessageBadRequest{}
}

// WriteResponse to the client
func (o *SendMessageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SendMessageUnauthorizedCode is the HTTP code returned for type SendMessageUnauthorized
const SendMessageUnauthorizedCode int = 401

/*SendMessageUnauthorized Authorization failure.

swagger:response sendMessageUnauthorized
*/
type SendMessageUnauthorized struct {
}

// NewSendMessageUnauthorized creates SendMessageUnauthorized with default headers values
func NewSendMessageUnauthorized() *SendMessageUnauthorized {

	return &SendMessageUnauthorized{}
}

// WriteResponse to the client
func (o *SendMessageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SendMessageForbiddenCode is the HTTP code returned for type SendMessageForbidden
const SendMessageForbiddenCode int = 403

/*SendMessageForbidden Quota exceeded or message too large.

swagger:response sendMessageForbidden
*/
type SendMessageForbidden struct {
}

// NewSendMessageForbidden creates SendMessageForbidden with default headers values
func NewSendMessageForbidden() *SendMessageForbidden {

	return &SendMessageForbidden{}
}

// WriteResponse to the client
func (o *SendMessageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// SendMessageGoneCode is the HTTP code returned for type SendMessageGone
const SendMessageGoneCode int = 410

/*SendMessageGone Specified queue or topic does not exist.

swagger:response sendMessageGone
*/
type SendMessageGone struct {
}

// NewSendMessageGone creates SendMessageGone with default headers values
func NewSendMessageGone() *SendMessageGone {

	return &SendMessageGone{}
}

// WriteResponse to the client
func (o *SendMessageGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(410)
}

// SendMessageInternalServerErrorCode is the HTTP code returned for type SendMessageInternalServerError
const SendMessageInternalServerErrorCode int = 500

/*SendMessageInternalServerError Internal error.

swagger:response sendMessageInternalServerError
*/
type SendMessageInternalServerError struct {
}

// NewSendMessageInternalServerError creates SendMessageInternalServerError with default headers values
func NewSendMessageInternalServerError() *SendMessageInternalServerError {

	return &SendMessageInternalServerError{}
}

// WriteResponse to the client
func (o *SendMessageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
