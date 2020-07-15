package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/abdulmoeid7112/student-backend-portal"
	"github.com/abdulmoeid7112/student-backend-portal/gen/restapi/operations"
)

// Handler replaces swagger handler.
type Handler *operations.StudentBackendPortalAPI

// NewHandler overrides swagger api handlers.
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewStudentBackendPortalAPI(spec)

	// student handlers.
	handler.AddStudentHandler = NewAddStudent(rt)
	handler.UpdateStudentHandler = NewUpdateStudent(rt)
	handler.DeleteStudentHandler = NewDeleteStudent(rt)
	handler.GetStudentHandler = NewGetStudent(rt)
	handler.ListStudentsHandler = NewListStudent(rt)

	return handler
}

