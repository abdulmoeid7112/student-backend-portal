// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Student student
//
// swagger:model student
type Student struct {

	// age
	Age int64 `json:"Age,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// level
	Level string `json:"Level,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// phone
	Phone string `json:"Phone,omitempty"`
}

// Validate validates this student
func (m *Student) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Student) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Student) UnmarshalBinary(b []byte) error {
	var res Student
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
