package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/abdulmoeid7112/student-backend-portal/config"
	"github.com/abdulmoeid7112/student-backend-portal/db"
	"github.com/abdulmoeid7112/student-backend-portal/models"
)

type client struct {
	db *sqlx.DB
}

func init() {
	db.Register("mysql", NewMysqlClient)
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

//NewMysqlClient to connect DB.
func NewMysqlClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	return &client{db: cli}, nil
}

func (m client) AddStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "id is not empty", nil
	}

	student.ID = guuid.NewV4().String()
	names := student.Names()
	if _, err := m.db.NamedExec(fmt.Sprintf(`INSERT INTO Student (%s) VALUES(%s)`, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")),
		student); err != nil {
		return "", errors.Wrap(err, "failed to add student")
	}
	return student.ID, nil
}

func (m client) UpdateStudent(id string, student *models.Student) error {
	names := student.Names()
	if _, err := m.db.NamedExec(fmt.Sprintf(`UPDATE Student SET %s WHERE id = '%s'`, strings.Join(mkPlaceHolder(names, "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ","), id), student); err != nil {
		return errors.Wrap(err, "failed to update student")
	}
	return nil
}

func (m client) GetStudentByID(id string) (*models.Student, error) {
	var std models.Student
	if err := m.db.Get(&std, fmt.Sprintf(`SELECT * FROM Student WHERE id='%s'`, id)); err != nil {
		if sql.ErrNoRows != nil {
			return nil, errors.Wrap(err, "failed to fetch student....not found")
		}
		return nil, err
	}
	fmt.Printf("%v\n", std.Map())
	return &std, nil
}

func (m client) ListStudent(filter map[string]interface{}, lim int64, off int64) ([]*models.Student, error) {
	var stdList []*models.Student
	if err := m.db.Select(&stdList, fmt.Sprintf("SELECT * FROM Student %s LIMIT %d,%d", mkFilter(filter), off, lim)); err != nil {
		if sql.ErrNoRows != nil {
			return stdList, errors.Wrap(err, "failed to list student....not found")
		}
		return stdList, err
	}
	log().Info(stdList)
	return stdList, nil
}

func (m client) RemoveStudentByID(id string) error {
	if _, err := m.db.Exec(fmt.Sprintf(`DELETE FROM Student WHERE id= '%s'`, id)); err != nil {
		return errors.Wrap(err, "failed to delete student")
	}
	return nil
}

func mkFilter(fil map[string]interface{}) string {
	if len(fil) < 1 {
		return ""
	}

	whereFilter := make([]string, 0)
	for key, val := range fil {
		whereFilter = append(whereFilter, fmt.Sprintf(" %s='%s' ", key, val))
	}

	return fmt.Sprintf("where %s", strings.Join(whereFilter, " AND "))
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}
	return ph
}
