// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExtUserObjectFlat user object flat
// swagger:model ExtUserObjectFlat
type ExtUserObjectFlat struct {

	// age
	Age int64 `json:"age,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user object flat
func (m *ExtUserObjectFlat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExtUserObjectFlat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtUserObjectFlat) UnmarshalBinary(b []byte) error {
	var res ExtUserObjectFlat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
