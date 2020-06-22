package db

import (
	"os"
	"reflect"
	"testing"
	"github.com/jmoiron/sqlx"
	"../models"
)

func Test_mysqlStd_AddStudent(t *testing.T) {

	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "student_portal")
	os.Setenv("DB_USER", "root")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		student *models.Student
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add student in db",
			args: args{student: &models.Student{
				Name:  "Usman",
				Age:   "21",
				Level: "Bachelor",
				Phone: "030051278936",
				Type:  "Student",
			}},
			wantErr: false,
		},
		{
			name: "fail - add invalid student in db",
			args: args{student: &models.Student{
				ID:    "1234",
				Name:  "Moeid",
				Age:   "21",
				Level: "Bachelor",
				Phone: "030051278936",
				Type:  "Employee"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMysqlfactory()
			_, err := m.AddStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mysqlStd_GetStudentByID(t *testing.T) {

	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "student_portal")
	os.Setenv("DB_USER", "root")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		id string
	}
	student := &models.Student{
		ID:    "1234",
		Name:  "Moeid",
		Age:   "21",
		Level: "Bachelor",
		Phone: "030051278936",
		Type:  "Employee"}
	tests := []struct {
		name    string
		args    args
		want    *models.Student
		wantErr bool
	}{
		{
			name:    "success - get student by ID",
			args:    args{id: "1234"},
			want:    student,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMysqlfactory()
			got, err := m.GetStudentByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStd_GetStudentEmployee(t *testing.T) {

	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "student_portal")
	os.Setenv("DB_USER", "root")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		stype string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - get students by Type",
			args:    args{stype: "Employee"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMysqlfactory()
			_, err := m.GetStudentEmployee(tt.args.stype)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudentEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mysqlStd_ListStudent(t *testing.T) {

	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "student_portal")
	os.Setenv("DB_USER", "root")

	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success - list all students",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMysqlfactory()
			_, err := m.ListStudent()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mysqlStd_UpdateStudent(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "student_portal")
	os.Setenv("DB_USER", "root")

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		student *models.Student
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - update records",
			args: args{student: &models.Student{
				ID:    "1234",
				Name:  "Moeid",
				Age:   "21",
				Level: "Bachelor",
				Phone: "030051278936",
				Type:  "Employee"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMysqlfactory()
			err := m.UpdateStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

