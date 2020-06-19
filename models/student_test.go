package models

import (
	"reflect"
	"testing"
)

func TestStudent_Map(t *testing.T) {
	type fields struct {
		Name  string
		Age   int32
		Level string
		Phone string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success",
			fields: fields{
				Name:  "Moeid",
				Age:   21,
				Level: "Bachelor",
				Phone: "030051278936",
			},
			want: map[string]interface{}{
				"Name":  "Moeid",
				"Age":   21,
				"Level": "Bachelor",
				"Phone": "030051278936",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
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
		Name  string
		Age   int32
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
				Name:  "Moeid",
				Age:   21,
				Level: "Bachelor",
				Phone: "030051278936",
			},
			want: []string{"Name", "Age", "Level", "Phone"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
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

