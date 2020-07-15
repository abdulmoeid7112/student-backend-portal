// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/abdulmoeid7112/student-backend-portal/gen/models"
)

// AddStudentReader is a Reader for the AddStudent structure.
type AddStudentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddStudentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddStudentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewAddStudentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddStudentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddStudentCreated creates a AddStudentCreated with default headers values
func NewAddStudentCreated() *AddStudentCreated {
	return &AddStudentCreated{}
}

/*AddStudentCreated handles this case with default header values.

student added
*/
type AddStudentCreated struct {
	Payload *models.Student
}

func (o *AddStudentCreated) Error() string {
	return fmt.Sprintf("[POST /student][%d] addStudentCreated  %+v", 201, o.Payload)
}

func (o *AddStudentCreated) GetPayload() *models.Student {
	return o.Payload
}

func (o *AddStudentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Student)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddStudentConflict creates a AddStudentConflict with default headers values
func NewAddStudentConflict() *AddStudentConflict {
	return &AddStudentConflict{}
}

/*AddStudentConflict handles this case with default header values.

student already exists
*/
type AddStudentConflict struct {
}

func (o *AddStudentConflict) Error() string {
	return fmt.Sprintf("[POST /student][%d] addStudentConflict ", 409)
}

func (o *AddStudentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddStudentInternalServerError creates a AddStudentInternalServerError with default headers values
func NewAddStudentInternalServerError() *AddStudentInternalServerError {
	return &AddStudentInternalServerError{}
}

/*AddStudentInternalServerError handles this case with default header values.

internal server error
*/
type AddStudentInternalServerError struct {
}

func (o *AddStudentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /student][%d] addStudentInternalServerError ", 500)
}

func (o *AddStudentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
