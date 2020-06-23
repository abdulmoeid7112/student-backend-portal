package models

import (
	"github.com/fatih/structs"
)

type Student struct {
	ID    string `json:"id" structs:"id" db:"id"`
	Name  string `json:"name" structs:"name" db:"name"`
	Age   int    `json:"age" structs:"age" db:"age"`
	Level string `json:"level" structs:"level" db:"level"`
	Phone string `json:"phone" structs:"phone" db:"phone"`
}

func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

func (s *Student) Names() []string {
	fields := structs.Fields(s)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}
