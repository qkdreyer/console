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
	"github.com/go-openapi/validate"
)

// SetConfigRequest set config request
//
// swagger:model setConfigRequest
type SetConfigRequest struct {

	// Used if configuration is an event notification's target
	ArnResourceID string `json:"arn_resource_id,omitempty"`

	// key values
	// Required: true
	// Min Items: 1
	KeyValues []*ConfigurationKV `json:"key_values"`
}

// Validate validates this set config request
func (m *SetConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetConfigRequest) validateKeyValues(formats strfmt.Registry) error {

	if err := validate.Required("key_values", "body", m.KeyValues); err != nil {
		return err
	}

	iKeyValuesSize := int64(len(m.KeyValues))

	if err := validate.MinItems("key_values", "body", iKeyValuesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.KeyValues); i++ {
		if swag.IsZero(m.KeyValues[i]) { // not required
			continue
		}

		if m.KeyValues[i] != nil {
			if err := m.KeyValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("key_values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("key_values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this set config request based on the context it is used
func (m *SetConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeyValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetConfigRequest) contextValidateKeyValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KeyValues); i++ {

		if m.KeyValues[i] != nil {
			if err := m.KeyValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("key_values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("key_values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetConfigRequest) UnmarshalBinary(b []byte) error {
	var res SetConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
