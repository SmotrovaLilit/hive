// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Address Address is a physical mailing address for this user. Canonical type values
// of "work", "home", and "other".  This attribute is a complex type
// with the following sub-attributes.
// swagger:model Address
type Address struct {

	// City: The city of the address.
	City string `json:"city,omitempty"`

	// Country: The country of the address.
	Country string `json:"country,omitempty"`

	// CountryCode: The [ISO 3166-1
	// alpha-2](http://www.iso.org/iso/country_codes.htm) country
	// code of the address.
	CountryCode string `json:"countryCode,omitempty"`

	// ExtendedAddress: The extended address of the address; for example,
	// the apartment number.
	ExtendedAddress string `json:"extendedAddress,omitempty"`

	// FormattedType: The read-only type of the address translated and
	// formatted in the viewer's
	// account locale or the `Accept-Language` HTTP header locale.
	FormattedType string `json:"formattedType,omitempty"`

	// FormattedValue: The unstructured value of the address. If this is not
	// set by the user it
	// will be automatically constructed from structured values.
	FormattedValue string `json:"formattedValue,omitempty"`

	// PoBox: The P.O. box of the address.
	PoBox string `json:"poBox,omitempty"`

	// PostalCode: The postal code of the address.
	PostalCode string `json:"postalCode,omitempty"`

	// Primary: True if the field is the primary field; false if the field
	// is a secondary
	// field.
	Primary bool `json:"primary,omitempty"`

	// Region: The region of the address; for example, the state or
	// province.
	Region string `json:"region,omitempty"`

	// StreetAddress: The street address.
	StreetAddress string `json:"streetAddress,omitempty"`

	// Type: The type of the address. The type can be custom or one of these
	// predefined values:
	//
	// `home`
	// `work`
	// `other`
	Type string `json:"type,omitempty"`

	// Verified: True if the field is verified; false if the field is
	// unverified. A
	// verified field is typically a name, email address, phone number,
	// or
	// website that has been confirmed to be owned by the person.
	Verified bool `json:"verified,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
