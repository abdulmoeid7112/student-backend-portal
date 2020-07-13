package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi/operations"
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

// NewAddStudent handles request for saving student.
func NewAddStudent(rt *runtime.Runtime) operations.AddStudentHandler {
	return &addStudent{rt: rt}
}

type addStudent struct {
	rt *runtime.Runtime
}

// Handle the add student request.
func (d *addStudent) Handle(params operations.AddStudentParams) middleware.Responder {
	student := models.Student{
		Name:  params.Student.Name,
		Age:   int(params.Student.Age),
		Level: params.Student.Level,
		Phone: params.Student.Phone,
	}

	id, err := d.rt.Service().AddStudent(&student)
	if err != nil {
		log().Errorf("failed to add student: %s", err)
		return operations.NewAddStudentConflict()
	}

	params.Student.ID = id
	return operations.NewAddStudentCreated().WithPayload(params.Student)
}

