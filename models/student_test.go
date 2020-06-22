package models

import (
	"reflect"
	"testing"
)

func TestStudent_Map(t *testing.T) {
	type fields struct {
		ID    string
		Name  string
		Age   int
		Level string
		Phone string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success-student",
			fields: fields{
				ID:    "1234",
				Name:  "Moeid",
				Age:   21,
				Level: "Bachelor",
				Phone: "030051278936",
			},
			want: map[string]interface{}{
				"id":    "1234",
				"name":  "Moeid",
				"age":   21,
				"level": "Bachelor",
				"phone": "030051278936",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Level: tt.fields.Level,
				Phone: tt.fields.Phone,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_Names(t *testing.T) {
	type fields struct {
		ID    string
		Name  string
		Age   int
		Level string
		Phone string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				ID:    "1234",
				Name:  "Moeid",
				Age:   21,
				Level: "Bachelor",
				Phone: "030051278936",
			},
			want: []string{"id", "name", "age", "level", "phone"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Level: tt.fields.Level,
				Phone: tt.fields.Phone,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

