package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi/operations"
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

// NewUpdateStudent function is for edit student.
func NewUpdateStudent(rt *runtime.Runtime) operations.UpdateStudentHandler {
	return &updateStudent{rt: rt}
}

type updateStudent struct {
	rt *runtime.Runtime
}

// Handle the put student request.
func (d *updateStudent) Handle(params operations.UpdateStudentParams) middleware.Responder {
	_, err := d.rt.Service().GetStudent(params.ID)
	if err != nil {
		return operations.NewGetStudentNotFound()
	}
	student := models.Student{
		Name:  params.Student.Name,
		Age:   int(params.Student.Age),
		Level: params.Student.Level,
		Phone: params.Student.Phone,
	}
	if err := d.rt.Service().UpdateStudent(params.ID, &student); err != nil {
		return operations.NewUpdateStudentInternalServerError()
	}

	return operations.NewUpdateStudentOK()
}

