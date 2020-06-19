package models

import (
	"github.com/fatih/structs"
)

type Student struct {
	Name  string `json:"name"`
	Age   int32  `json:"age"`
	Level string `json:"level"`
	Phone string `json:"phone"`
}

func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

func (s *Student) Names() []string {
	return structs.Names(s)
}

