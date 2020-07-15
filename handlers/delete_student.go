package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	domainErr "github.com/abdulmoeid7112/student-backend-portal/errors"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi/operations"
)

// NewDeleteStudent function will delete the student.
func NewDeleteStudent(rt *runtime.Runtime) operations.DeleteStudentHandler {
	return &deleteStudent{rt: rt}
}

type deleteStudent struct {
	rt *runtime.Runtime
}

// Handle the delete entry request.
func (d *deleteStudent) Handle(params operations.DeleteStudentParams) middleware.Responder {
	if err := d.rt.Service().DeleteStudent(params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetStudentNotFound()
		default:
			return operations.NewDeleteStudentInternalServerError()
		}
	}
	return operations.NewDeleteStudentNoContent()
}

