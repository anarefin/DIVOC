// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificationRequest certification request
//
// swagger:model CertificationRequest
type CertificationRequest struct {

	// facility
	Facility *CertificationRequestFacility `json:"facility,omitempty"`

	// meta
	Meta interface{} `json:"meta,omitempty"`

	// pre enrollment code
	PreEnrollmentCode string `json:"preEnrollmentCode,omitempty"`

	// recipient
	Recipient *CertificationRequestRecipient `json:"recipient,omitempty"`

	// vaccination
	Vaccination *CertificationRequestVaccination `json:"vaccination,omitempty"`

	// vaccinator
	Vaccinator *CertificationRequestVaccinator `json:"vaccinator,omitempty"`
}

// Validate validates this certification request
func (m *CertificationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaccinator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequest) validateFacility(formats strfmt.Registry) error {

	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if m.Facility != nil {
		if err := m.Facility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequest) validateRecipient(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipient) { // not required
		return nil
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequest) validateVaccination(formats strfmt.Registry) error {

	if swag.IsZero(m.Vaccination) { // not required
		return nil
	}

	if m.Vaccination != nil {
		if err := m.Vaccination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vaccination")
			}
			return err
		}
	}

	return nil
}

func (m *CertificationRequest) validateVaccinator(formats strfmt.Registry) error {

	if swag.IsZero(m.Vaccinator) { // not required
		return nil
	}

	if m.Vaccinator != nil {
		if err := m.Vaccinator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vaccinator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequest) UnmarshalBinary(b []byte) error {
	var res CertificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestFacility certification request facility
//
// swagger:model CertificationRequestFacility
type CertificationRequestFacility struct {

	// address
	Address *CertificationRequestFacilityAddress `json:"address,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this certification request facility
func (m *CertificationRequestFacility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestFacility) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("facility" + "." + "address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestFacility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestFacility) UnmarshalBinary(b []byte) error {
	var res CertificationRequestFacility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestFacilityAddress certification request facility address
//
// swagger:model CertificationRequestFacilityAddress
type CertificationRequestFacilityAddress struct {

	// address line1
	AddressLine1 string `json:"addressLine1,omitempty"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// district
	District string `json:"district,omitempty"`

	// pincode
	Pincode int64 `json:"pincode,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this certification request facility address
func (m *CertificationRequestFacilityAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestFacilityAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestFacilityAddress) UnmarshalBinary(b []byte) error {
	var res CertificationRequestFacilityAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestRecipient certification request recipient
//
// swagger:model CertificationRequestRecipient
type CertificationRequestRecipient struct {

	// age
	Age string `json:"age,omitempty"`

	// contact
	Contact []string `json:"contact"`

	// dob
	// Format: date
	Dob strfmt.Date `json:"dob,omitempty"`

	// gender
	Gender string `json:"gender,omitempty"`

	// identity
	Identity string `json:"identity,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nationality
	Nationality string `json:"nationality,omitempty"`
}

// Validate validates this certification request recipient
func (m *CertificationRequestRecipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestRecipient) validateDob(formats strfmt.Registry) error {

	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := validate.FormatOf("recipient"+"."+"dob", "body", "date", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestRecipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestRecipient) UnmarshalBinary(b []byte) error {
	var res CertificationRequestRecipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestVaccination certification request vaccination
//
// swagger:model CertificationRequestVaccination
type CertificationRequestVaccination struct {

	// batch
	Batch string `json:"batch,omitempty"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// effective start
	// Format: date
	EffectiveStart strfmt.Date `json:"effectiveStart,omitempty"`

	// effective until
	// Format: date
	EffectiveUntil strfmt.Date `json:"effectiveUntil,omitempty"`

	// manufacturer
	Manufacturer string `json:"manufacturer,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this certification request vaccination
func (m *CertificationRequestVaccination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificationRequestVaccination) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("vaccination"+"."+"date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateEffectiveStart(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveStart) { // not required
		return nil
	}

	if err := validate.FormatOf("vaccination"+"."+"effectiveStart", "body", "date", m.EffectiveStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CertificationRequestVaccination) validateEffectiveUntil(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("vaccination"+"."+"effectiveUntil", "body", "date", m.EffectiveUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestVaccination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestVaccination) UnmarshalBinary(b []byte) error {
	var res CertificationRequestVaccination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CertificationRequestVaccinator certification request vaccinator
//
// swagger:model CertificationRequestVaccinator
type CertificationRequestVaccinator struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this certification request vaccinator
func (m *CertificationRequestVaccinator) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificationRequestVaccinator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationRequestVaccinator) UnmarshalBinary(b []byte) error {
	var res CertificationRequestVaccinator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
