package service

import (
	"github.com/abdulmoeid7112/student-backend-portal/db"
)

// Service initializes our database instance.
type Service struct {
	db db.DataStore
}

// NewService creates a connection to our database.
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}

