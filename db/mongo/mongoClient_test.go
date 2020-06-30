package mongo

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/abdulmoeid7112/student-backend-portal/db"
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

func Test_mongoStd_AddStudent(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "student-portal-mongo-db")
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
				Name:  "Moeid",
				Age:   22,
				Level: "Bachelor",
				Phone: "03005124896",
			}},
			wantErr: false,
		},
		{
			name: "fail - add invalid student in db",
			args: args{student: &models.Student{
				ID:    "1234",
				Name:  "Moeid",
				Age:   21,
				Level: "Bachelor",
				Phone: "030051278936",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewMongoClient(db.Option{})
			_, err := m.AddStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mongoStd_UpdateStudent(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "student-portal-mongo-db")

	m, _ := NewMongoClient(db.Option{})
	student := &models.Student{
		Name:  "Ali",
		Age:   22,
		Level: "Bachelor",
		Phone: "03005124896",
	}
	_, _ = m.AddStudent(student)
	type args struct {
		id      string
		student *models.Student
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - update records",
			args: args{id: student.ID,
				student: &models.Student{
					Name:  "Ateeq",
					Age:   21,
					Level: "Master",
					Phone: "03005412263",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.UpdateStudent(tt.args.id, tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mongoStd_GetStudentByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "student-portal-mongo-db")

	m, _ := NewMongoClient(db.Option{})
	student := &models.Student{
		Name:  "Ali",
		Age:   22,
		Level: "Bachelor",
		Phone: "03005124896",
	}
	_, _ = m.AddStudent(student)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Student
		wantErr bool
	}{
		{
			name:    "success - get student by ID",
			args:    args{id: student.ID},
			want:    student,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetStudentByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if cmp.Equal(got, tt.want) {
				log().Info(got)
			} else {
				t.Errorf("GetStudentByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mongoStd_RemoveStudentByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "student-portal-mongo-db")

	m, _ := NewMongoClient(db.Option{})
	student := &models.Student{
		Name:  "Usman",
		Age:   22,
		Level: "Bachelor",
		Phone: "03005124896",
	}
	_, _ = m.AddStudent(student)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Student
		wantErr bool
	}{
		{
			name:    "success - delete student by ID",
			args:    args{id: student.ID},
			want:    student,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.RemoveStudentByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveStudentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_mongoStd_ListStudent(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "student-portal-mongo-db")
	m, _ := NewMongoClient(db.Option{})
	student1 := &models.Student{
		Name:  "Sohaib",
		Age:   25,
		Level: "Master",
		Phone: "03024563789",
	}
	student2 := &models.Student{
		Name:  "Usama",
		Age:   28,
		Level: "PHD",
		Phone: "03217895641",
	}
	_, _ = m.AddStudent(student1)
	_, _ = m.AddStudent(student2)
	type args struct {
		filter map[string]interface{}
		lim    int64
		off    int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Student
		wantErr bool
	}{
		{
			name:    "success- list all students with name",
			args:    args{filter: map[string]interface{}{"name": "Sohaib"}, lim: 1, off: 0},
			want:    []*models.Student{student1},
			wantErr: false,
		},
		{
			name:    "fail- list all students with name",
			args:    args{filter: map[string]interface{}{"name": "Aftab"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all students with name level",
			args:    args{filter: map[string]interface{}{"name": "Sohaib","level": "Master"}, lim: 1, off: 0},
			want:    []*models.Student{student1},
			wantErr: false,
		},
		{
			name:    "fail- list all students with name level",
			args:    args{filter: map[string]interface{}{"name": "Aftab","level": "Master"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all students with phone ",
			args:    args{filter: map[string]interface{}{"phone": "03217895641"}, lim: 1, off: 0},
			want:    []*models.Student{student2},
			wantErr: false,
		},
		{
			name:    "fail- list all students with phone",
			args:    args{filter: map[string]interface{}{"phone": "03214569853"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "success- list all students with phone age",
			args:    args{filter: map[string]interface{}{"age": 28, "phone": "03217895641"}, lim: 1, off: 0},
			want:    []*models.Student{student2},
			wantErr: false,
		},
		{
			name:    "fail- list all students with name Hameed age 22",
			args:    args{filter: map[string]interface{}{"age": 22, "phone": "031458795149"}, lim: 1, off: 0},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.ListStudent(tt.args.filter, tt.args.lim, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ListStudent() got = %v, want = %v,diff = %s, error = %v", got, tt.want, diff, err)
			}else {
				log().Info(got)
			}
		})
	}
}
