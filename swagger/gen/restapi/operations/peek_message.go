// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PeekMessageHandlerFunc turns a function with the right signature into a peek message handler
type PeekMessageHandlerFunc func(PeekMessageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PeekMessageHandlerFunc) Handle(params PeekMessageParams) middleware.Responder {
	return fn(params)
}

// PeekMessageHandler interface for that can handle valid peek message params
type PeekMessageHandler interface {
	Handle(PeekMessageParams) middleware.Responder
}

// NewPeekMessage creates a new http.Handler for the peek message operation
func NewPeekMessage(ctx *middleware.Context, handler PeekMessageHandler) *PeekMessage {
	return &PeekMessage{Context: ctx, Handler: handler}
}

/* PeekMessage swagger:route POST /{queueName}/messages/head peekMessage

PeekMessage peek message API

*/
type PeekMessage struct {
	Context *middleware.Context
	Handler PeekMessageHandler
}

func (o *PeekMessage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPeekMessageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
