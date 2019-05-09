// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixVendorVerifyInfo openpitrix vendor verify info
// swagger:model openpitrixVendorVerifyInfo
type OpenpitrixVendorVerifyInfo struct {

	// approver who approve the vendor verify
	Approver string `json:"approver,omitempty"`

	// authorizer email eg.***@yunify.com
	AuthorizerEmail string `json:"authorizer_email,omitempty"`

	// authorizer name
	AuthorizerName string `json:"authorizer_name,omitempty"`

	// authorizer phone, string of 11 digit
	AuthorizerPhone string `json:"authorizer_phone,omitempty"`

	// name of bank account
	BankAccountName string `json:"bank_account_name,omitempty"`

	// number of bank account
	BankAccountNumber string `json:"bank_account_number,omitempty"`

	// bank name
	BankName string `json:"bank_name,omitempty"`

	// company name
	CompanyName string `json:"company_name,omitempty"`

	// company profile
	CompanyProfile string `json:"company_profile,omitempty"`

	// company website
	CompanyWebsite string `json:"company_website,omitempty"`

	// owner who own the vendor verify
	Owner string `json:"owner,omitempty"`

	// owner path, concat string group_path:user_id
	OwnerPath string `json:"owner_path,omitempty"`

	// reject message
	RejectMessage string `json:"reject_message,omitempty"`

	// status eg.[draft|submitted|passed|rejected|suspended|in-review|new]
	Status string `json:"status,omitempty"`

	// record status changed time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// submit time of vendor verify
	SubmitTime strfmt.DateTime `json:"submit_time,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix vendor verify info
func (m *OpenpitrixVendorVerifyInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixVendorVerifyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixVendorVerifyInfo) UnmarshalBinary(b []byte) error {
	var res OpenpitrixVendorVerifyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
