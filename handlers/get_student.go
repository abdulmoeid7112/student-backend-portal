package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	domainErr "github.com/abdulmoeid7112/student-backend-portal/errors"
	"github.com/abdulmoeid7112/student-backend-portal/gen/models"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi/operations"
)

// NewGetStudent handles a request for retrieving student.
func NewGetStudent(rt *runtime.Runtime) operations.GetStudentHandler {
	return &getStudent{rt: rt}
}

type getStudent struct {
	rt *runtime.Runtime
}

// Handle the get student request.
func (d *getStudent) Handle(params operations.GetStudentParams) middleware.Responder {
	student, err := d.rt.Service().GetStudent(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetStudentNotFound()
		default:
			return operations.NewGetStudentInternalServerError()
		}
	}

	return operations.NewGetStudentOK().WithPayload(&models.Student{
		Age:   int64(student.Age),
		Level: student.Level,
		ID:    student.ID,
		Name:  student.Name,
		Phone: student.Phone,
	})
}

