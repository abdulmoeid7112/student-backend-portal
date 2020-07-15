package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	domainErr "github.com/abdulmoeid7112/student-backend-portal/errors"
	"github.com/abdulmoeid7112/student-backend-portal/gen/models"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi/operations"
)

// NewListStudent handles a request for retrieving students.
func NewListStudent(rt *runtime.Runtime) operations.ListStudentsHandler {
	return &listStudents{rt: rt}
}

type listStudents struct {
	rt *runtime.Runtime
}

// Handle the list employees request.
func (e *listStudents) Handle(params operations.ListStudentsParams) middleware.Responder {
	filter := make(map[string]interface{})
	if params.Name!=nil {
		filter["name"] = swag.StringValue(params.Name)
	}
	if params.Age!=nil {
		filter["age"] = swag.Int64Value(params.Age)
	}
	if params.Level!=nil {
		filter["level"] = swag.StringValue(params.Level)
	}
	if params.Phone!=nil {
		filter["phone"] = swag.StringValue(params.Phone)
	}

	students, err := e.rt.Service().ListStudents(filter, swag.Int64Value(params.Limit), swag.Int64Value(params.Offset))
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetStudentNotFound()
		default:
			return operations.NewUpdateStudentInternalServerError()
		}
	}

	stdList := make([]*models.Student, 0)
	for _,student := range students {
		stdList = append(stdList, toStudentGen(student))
	}

	return operations.NewListStudentsOK().WithPayload(stdList)
}

