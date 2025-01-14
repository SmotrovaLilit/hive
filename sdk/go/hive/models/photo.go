// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Photo photo
// swagger:model Photo
type Photo struct {

	// Primary: True if the field is the primary field; false if the field
	// is a secondary
	// field.
	Primary bool `json:"primary,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`

	// Verified: True if the field is verified; false if the field is
	// unverified. A
	// verified field is typically a name, email address, phone number,
	// or
	// website that has been confirmed to be owned by the person.
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this photo
func (m *Photo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Photo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Photo) UnmarshalBinary(b []byte) error {
	var res Photo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
