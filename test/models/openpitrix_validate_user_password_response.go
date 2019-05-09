// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixValidateUserPasswordResponse openpitrix validate user password response
// swagger:model openpitrixValidateUserPasswordResponse
type OpenpitrixValidateUserPasswordResponse struct {

	// validate password ok or not
	Validated bool `json:"validated,omitempty"`
}

// Validate validates this openpitrix validate user password response
func (m *OpenpitrixValidateUserPasswordResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixValidateUserPasswordResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixValidateUserPasswordResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixValidateUserPasswordResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
