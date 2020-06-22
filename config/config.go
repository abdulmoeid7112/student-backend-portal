package config

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func GetMySQLDB() string {
	dbUser := "root"
	dbPass := ""
	dbName := "student_portal"
	dbHost := "localhost"
	dbPort := "3306"
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", dbHost, dbPort)
	cfg.DBName = dbName
	cfg.ParseTime = true
	cfg.User = dbUser
	cfg.Passwd = dbPass

	return cfg.FormatDSN()
}
