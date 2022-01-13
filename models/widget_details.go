// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WidgetDetails widget details
//
// swagger:model widgetDetails
type WidgetDetails struct {

	// id
	ID int32 `json:"id,omitempty"`

	// options
	Options *WidgetDetailsOptions `json:"options,omitempty"`

	// targets
	Targets []*ResultTarget `json:"targets"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this widget details
func (m *WidgetDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WidgetDetails) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *WidgetDetails) validateTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.Targets) { // not required
		return nil
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this widget details based on the context it is used
func (m *WidgetDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WidgetDetails) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.Options != nil {
		if err := m.Options.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *WidgetDetails) contextValidateTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Targets); i++ {

		if m.Targets[i] != nil {
			if err := m.Targets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WidgetDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WidgetDetails) UnmarshalBinary(b []byte) error {
	var res WidgetDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WidgetDetailsOptions widget details options
//
// swagger:model WidgetDetailsOptions
type WidgetDetailsOptions struct {

	// reduce options
	ReduceOptions *WidgetDetailsOptionsReduceOptions `json:"reduceOptions,omitempty"`
}

// Validate validates this widget details options
func (m *WidgetDetailsOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReduceOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WidgetDetailsOptions) validateReduceOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ReduceOptions) { // not required
		return nil
	}

	if m.ReduceOptions != nil {
		if err := m.ReduceOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options" + "." + "reduceOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options" + "." + "reduceOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this widget details options based on the context it is used
func (m *WidgetDetailsOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReduceOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WidgetDetailsOptions) contextValidateReduceOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ReduceOptions != nil {
		if err := m.ReduceOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options" + "." + "reduceOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options" + "." + "reduceOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WidgetDetailsOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WidgetDetailsOptions) UnmarshalBinary(b []byte) error {
	var res WidgetDetailsOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WidgetDetailsOptionsReduceOptions widget details options reduce options
//
// swagger:model WidgetDetailsOptionsReduceOptions
type WidgetDetailsOptionsReduceOptions struct {

	// calcs
	Calcs []string `json:"calcs"`
}

// Validate validates this widget details options reduce options
func (m *WidgetDetailsOptionsReduceOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this widget details options reduce options based on context it is used
func (m *WidgetDetailsOptionsReduceOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WidgetDetailsOptionsReduceOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WidgetDetailsOptionsReduceOptions) UnmarshalBinary(b []byte) error {
	var res WidgetDetailsOptionsReduceOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
