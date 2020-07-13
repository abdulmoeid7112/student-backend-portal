package service

import (
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

// AddStudent adds or update student into database.
func (s *Service) AddStudent(student *models.Student) (string, error) {
	return s.db.AddStudent(student)
}

// UpdateStudent updates student record in database.
func (s *Service) UpdateStudent(id string, student *models.Student) error {
	return s.db.UpdateStudent(id, student)
}

// GetStudent retrieve student from database.
func (s *Service) GetStudent(id string) (*models.Student, error) {
	return s.db.GetStudentByID(id)
}

// DeleteStudent deletes student from database.
func (s *Service) DeleteStudent(id string) error {
	return s.db.RemoveStudentByID(id)
}

// ListStudents retrieve all the students from database.
func (s *Service) ListStudents(filter map[string]interface{}, lim int64, off int64) ([]*models.Student, error) {
	return s.db.ListStudent(filter, lim, off)
}

