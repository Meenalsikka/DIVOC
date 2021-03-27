// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Facility facility
//
// swagger:model Facility
type Facility struct {

	// Address
	Address *Address `json:"address,omitempty"`

	// admins
	Admins []*FacilityAdmin `json:"admins"`

	// Average Rating
	//
	// Average Rating of Facility 0 to 5, 0 for no rating.
	AverageRating float64 `json:"averageRating,omitempty"`

	// Category
	// Enum: [GOVT PRIVATE]
	Category string `json:"category,omitempty"`

	// Contact number
	Contact string `json:"contact,omitempty"`

	// Facility Email
	Email string `json:"email,omitempty"`

	// Facility Code
	FacilityCode string `json:"facilityCode,omitempty"`

	// Facility Name
	FacilityName string `json:"facilityName,omitempty"`

	// Geo Location
	GeoLocation string `json:"geoLocation,omitempty"`

	// Operating hours end of day
	OperatingHourEnd string `json:"operatingHourEnd,omitempty"`

	// Operating hours start of day
	OperatingHourStart string `json:"operatingHourStart,omitempty"`

	// programs
	Programs []*FacilityProgramsItems0 `json:"programs"`

	// stamp
	Stamp string `json:"stamp,omitempty"`

	// Status of Facility
	// Enum: [Active Inactive Blocked]
	Status string `json:"status,omitempty"`

	// Type of Facility
	// Enum: [Fixed location Mobile Both]
	Type string `json:"type,omitempty"`

	// Website URL
	WebsiteURL string `json:"websiteUrl,omitempty"`
}

// Validate validates this facility
func (m *Facility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdmins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrograms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Facility) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *Facility) validateAdmins(formats strfmt.Registry) error {

	if swag.IsZero(m.Admins) { // not required
		return nil
	}

	for i := 0; i < len(m.Admins); i++ {
		if swag.IsZero(m.Admins[i]) { // not required
			continue
		}

		if m.Admins[i] != nil {
			if err := m.Admins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("admins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var facilityTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GOVT","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		facilityTypeCategoryPropEnum = append(facilityTypeCategoryPropEnum, v)
	}
}

const (

	// FacilityCategoryGOVT captures enum value "GOVT"
	FacilityCategoryGOVT string = "GOVT"

	// FacilityCategoryPRIVATE captures enum value "PRIVATE"
	FacilityCategoryPRIVATE string = "PRIVATE"
)

// prop value enum
func (m *Facility) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, facilityTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Facility) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *Facility) validatePrograms(formats strfmt.Registry) error {

	if swag.IsZero(m.Programs) { // not required
		return nil
	}

	for i := 0; i < len(m.Programs); i++ {
		if swag.IsZero(m.Programs[i]) { // not required
			continue
		}

		if m.Programs[i] != nil {
			if err := m.Programs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("programs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var facilityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive","Blocked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		facilityTypeStatusPropEnum = append(facilityTypeStatusPropEnum, v)
	}
}

const (

	// FacilityStatusActive captures enum value "Active"
	FacilityStatusActive string = "Active"

	// FacilityStatusInactive captures enum value "Inactive"
	FacilityStatusInactive string = "Inactive"

	// FacilityStatusBlocked captures enum value "Blocked"
	FacilityStatusBlocked string = "Blocked"
)

// prop value enum
func (m *Facility) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, facilityTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Facility) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var facilityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Fixed location","Mobile","Both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		facilityTypeTypePropEnum = append(facilityTypeTypePropEnum, v)
	}
}

const (

	// FacilityTypeFixedLocation captures enum value "Fixed location"
	FacilityTypeFixedLocation string = "Fixed location"

	// FacilityTypeMobile captures enum value "Mobile"
	FacilityTypeMobile string = "Mobile"

	// FacilityTypeBoth captures enum value "Both"
	FacilityTypeBoth string = "Both"
)

// prop value enum
func (m *Facility) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, facilityTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Facility) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Facility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Facility) UnmarshalBinary(b []byte) error {
	var res Facility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FacilityProgramsItems0 facility programs items0
//
// swagger:model FacilityProgramsItems0
type FacilityProgramsItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// program Id
	ProgramID string `json:"programId,omitempty"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// rate updated at
	RateUpdatedAt string `json:"rateUpdatedAt,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status updated at
	StatusUpdatedAt string `json:"statusUpdatedAt,omitempty"`
}

// Validate validates this facility programs items0
func (m *FacilityProgramsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FacilityProgramsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityProgramsItems0) UnmarshalBinary(b []byte) error {
	var res FacilityProgramsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
