// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListStudentsHandlerFunc turns a function with the right signature into a list students handler
type ListStudentsHandlerFunc func(ListStudentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStudentsHandlerFunc) Handle(params ListStudentsParams) middleware.Responder {
	return fn(params)
}

// ListStudentsHandler interface for that can handle valid list students params
type ListStudentsHandler interface {
	Handle(ListStudentsParams) middleware.Responder
}

// NewListStudents creates a new http.Handler for the list students operation
func NewListStudents(ctx *middleware.Context, handler ListStudentsHandler) *ListStudents {
	return &ListStudents{Context: ctx, Handler: handler}
}

/*ListStudents swagger:route GET /student listStudents

get all the students

*/
type ListStudents struct {
	Context *middleware.Context
	Handler ListStudentsHandler
}

func (o *ListStudents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStudentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
