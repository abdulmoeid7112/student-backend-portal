package db

import (
	"../config"
	"../models"
	"database/sql"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
)

var UserNotFoundError = errors.New("User not found")
var FailedError = errors.New("failed to connect...")
var FailedAdd = errors.New("failed to add...")
var FailedUpdate = errors.New("failed to update...")
var FailedDelete = errors.New("failed to delete...")

type DataStore interface {
	AddStudent(student *models.Student) (string, error)
	UpdateStudent(student *models.Student) error
	GetStudentByID(id string) (*models.Student, error)
	ListStudent() (*[]models.Student, error)
	RemoveStudentByID(id string) error
}

type mysqlStd struct {
	db *sqlx.DB
}

func (m mysqlStd) AddStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "", errors.New("id is not empty")
	}
	student.ID = guuid.New().String()

	if _, err := m.db.NamedExec(fmt.Sprintf(`INSERT INTO 
		Student (id,name,age,level,phone) 
		VALUES(:id,:name,:age,:level,:phone)`),
		student); err != nil {
		return "", FailedAdd
	}
	return "", nil
}

func (m mysqlStd) UpdateStudent(student *models.Student) error {
	if _, err := m.db.NamedExec(fmt.Sprintf(`UPDATE Student SET 
	id=:id, age=:age, level=:level, phone=:phone
	WHERE name=:name`), student); err != nil {
		return FailedUpdate
	}
	return nil
}

func (m mysqlStd) GetStudentByID(id string) (*models.Student, error) {
	var std models.Student
	if err := m.db.Get(&std, fmt.Sprintf(`SELECT * FROM Student WHERE id='%s'`, id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, UserNotFoundError
		}
		return nil, err
	}
	fmt.Printf("%v\n", std.Map())
	return &std, nil
}

func (m mysqlStd) ListStudent() (*[]models.Student, error) {
	var std []models.Student
	if err := m.db.Select(&std, fmt.Sprint(`SELECT * FROM Student`)); err != nil {
		if err == sql.ErrNoRows {
			return nil, UserNotFoundError
		}
		return nil, err
	}
	for i := 0; i < len(std); i++ {
		fmt.Printf("%v\n", std[i].Map())
	}

	return &std, nil
}
func (m mysqlStd) RemoveStudentByID(id string) error {
	if _, err := m.db.Query(fmt.Sprintf(`DELETE FROM Student WHERE id= '%s'`, id)); err != nil {
		return FailedDelete
	}

	return nil
}

type DataStoreFactory func() (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
}

func NewMysqlfactory() (DataStore, error) {
	cli, err := sqlx.Connect("mysql", config.GetMySQLDB())
	if err != nil {
		return nil, FailedError
	}
	return &mysqlStd{db: cli}, nil
}

func init() {
	Register("mysql", NewMysqlfactory)
}

