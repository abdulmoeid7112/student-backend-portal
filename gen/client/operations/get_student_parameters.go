// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetStudentParams creates a new GetStudentParams object
// with the default values initialized.
func NewGetStudentParams() *GetStudentParams {
	var ()
	return &GetStudentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStudentParamsWithTimeout creates a new GetStudentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStudentParamsWithTimeout(timeout time.Duration) *GetStudentParams {
	var ()
	return &GetStudentParams{

		timeout: timeout,
	}
}

// NewGetStudentParamsWithContext creates a new GetStudentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStudentParamsWithContext(ctx context.Context) *GetStudentParams {
	var ()
	return &GetStudentParams{

		Context: ctx,
	}
}

// NewGetStudentParamsWithHTTPClient creates a new GetStudentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStudentParamsWithHTTPClient(client *http.Client) *GetStudentParams {
	var ()
	return &GetStudentParams{
		HTTPClient: client,
	}
}

/*GetStudentParams contains all the parameters to send to the API endpoint
for the get student operation typically these are written to a http.Request
*/
type GetStudentParams struct {

	/*ID
	  UUID of the student

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get student params
func (o *GetStudentParams) WithTimeout(timeout time.Duration) *GetStudentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get student params
func (o *GetStudentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get student params
func (o *GetStudentParams) WithContext(ctx context.Context) *GetStudentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get student params
func (o *GetStudentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get student params
func (o *GetStudentParams) WithHTTPClient(client *http.Client) *GetStudentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get student params
func (o *GetStudentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get student params
func (o *GetStudentParams) WithID(id string) *GetStudentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get student params
func (o *GetStudentParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetStudentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
