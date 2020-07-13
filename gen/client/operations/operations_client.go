// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteStudent(params *DeleteStudentParams) (*DeleteStudentNoContent, error)

	ListStudents(params *ListStudentsParams) (*ListStudentsOK, error)

	AddStudent(params *AddStudentParams) (*AddStudentCreated, error)

	GetStudent(params *GetStudentParams) (*GetStudentOK, error)

	UpdateStudent(params *UpdateStudentParams) (*UpdateStudentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteStudent delete student with id provided in url
*/
func (a *Client) DeleteStudent(params *DeleteStudentParams) (*DeleteStudentNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStudentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteStudent",
		Method:             "DELETE",
		PathPattern:        "/student/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteStudentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteStudent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListStudents get all the students
*/
func (a *Client) ListStudents(params *ListStudentsParams) (*ListStudentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListStudentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListStudents",
		Method:             "GET",
		PathPattern:        "/student",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListStudentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListStudentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListStudents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddStudent add student API
*/
func (a *Client) AddStudent(params *AddStudentParams) (*AddStudentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddStudentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addStudent",
		Method:             "POST",
		PathPattern:        "/student",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddStudentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addStudent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStudent return student based on UUID
*/
func (a *Client) GetStudent(params *GetStudentParams) (*GetStudentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStudentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStudent",
		Method:             "GET",
		PathPattern:        "/student/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStudentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStudent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateStudent update student API
*/
func (a *Client) UpdateStudent(params *UpdateStudentParams) (*UpdateStudentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStudentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateStudent",
		Method:             "PUT",
		PathPattern:        "/student/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateStudentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateStudentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateStudent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
