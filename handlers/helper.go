package handlers

import (
	genModels "github.com/abdulmoeid7112/student-backend-portal/gen/models"
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

func toStudentGen(student *models.Student) *genModels.Student{
	return &genModels.Student{
		ID: student.ID,
		Name:  student.Name,
		Age:   int64(student.Age),
		Level: student.Level,
		Phone: student.Phone,
	}
}

