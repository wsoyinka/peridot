// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginConfigLinux PluginConfigLinux plugin config linux
//
// swagger:model PluginConfigLinux
type PluginConfigLinux struct {

	// allow all devices
	// Required: true
	AllowAllDevices *bool `json:"AllowAllDevices"`

	// capabilities
	// Required: true
	Capabilities []string `json:"Capabilities"`

	// devices
	// Required: true
	Devices []*PluginDevice `json:"Devices"`
}

// Validate validates this plugin config linux
func (m *PluginConfigLinux) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowAllDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigLinux) validateAllowAllDevices(formats strfmt.Registry) error {

	if err := validate.Required("AllowAllDevices", "body", m.AllowAllDevices); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigLinux) validateCapabilities(formats strfmt.Registry) error {

	if err := validate.Required("Capabilities", "body", m.Capabilities); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigLinux) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("Devices", "body", m.Devices); err != nil {
		return err
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugin config linux based on the context it is used
func (m *PluginConfigLinux) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigLinux) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Devices); i++ {

		if m.Devices[i] != nil {
			if err := m.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigLinux) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigLinux) UnmarshalBinary(b []byte) error {
	var res PluginConfigLinux
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
