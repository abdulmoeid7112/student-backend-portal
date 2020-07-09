package models

import (
	"github.com/fatih/structs"
)

// Student structs.
type Student struct {
	ID    string `json:"id" structs:"id" bson:"_id" db:"id"`
	Name  string `json:"name" structs:"name" bson:"name" db:"name"`
	Age   int    `json:"age" structs:"age" bson:"age" db:"age"`
	Level string `json:"level" structs:"level" bson:"level" db:"level"`
	Phone string `json:"phone" structs:"phone" bson:"phone" db:"phone"`
}

// Map function returns map values.
func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names function returns field names.
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
