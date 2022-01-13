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

// IdpConfiguration idp configuration
//
// swagger:model idpConfiguration
type IdpConfiguration struct {

	// active directory
	ActiveDirectory *IdpConfigurationActiveDirectory `json:"active_directory,omitempty"`

	// keys
	Keys []*IdpConfigurationKeysItems0 `json:"keys"`

	// oidc
	Oidc *IdpConfigurationOidc `json:"oidc,omitempty"`
}

// Validate validates this idp configuration
func (m *IdpConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfiguration) validateActiveDirectory(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveDirectory) { // not required
		return nil
	}

	if m.ActiveDirectory != nil {
		if err := m.ActiveDirectory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_directory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("active_directory")
			}
			return err
		}
	}

	return nil
}

func (m *IdpConfiguration) validateKeys(formats strfmt.Registry) error {
	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	for i := 0; i < len(m.Keys); i++ {
		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {
			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdpConfiguration) validateOidc(formats strfmt.Registry) error {
	if swag.IsZero(m.Oidc) { // not required
		return nil
	}

	if m.Oidc != nil {
		if err := m.Oidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this idp configuration based on the context it is used
func (m *IdpConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActiveDirectory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOidc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfiguration) contextValidateActiveDirectory(ctx context.Context, formats strfmt.Registry) error {

	if m.ActiveDirectory != nil {
		if err := m.ActiveDirectory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_directory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("active_directory")
			}
			return err
		}
	}

	return nil
}

func (m *IdpConfiguration) contextValidateKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Keys); i++ {

		if m.Keys[i] != nil {
			if err := m.Keys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdpConfiguration) contextValidateOidc(ctx context.Context, formats strfmt.Registry) error {

	if m.Oidc != nil {
		if err := m.Oidc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfiguration) UnmarshalBinary(b []byte) error {
	var res IdpConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdpConfigurationActiveDirectory idp configuration active directory
//
// swagger:model IdpConfigurationActiveDirectory
type IdpConfigurationActiveDirectory struct {

	// group name attribute
	GroupNameAttribute string `json:"group_name_attribute,omitempty"`

	// group search base dn
	GroupSearchBaseDn string `json:"group_search_base_dn,omitempty"`

	// group search filter
	GroupSearchFilter string `json:"group_search_filter,omitempty"`

	// lookup bind dn
	LookupBindDn string `json:"lookup_bind_dn,omitempty"`

	// lookup bind password
	LookupBindPassword string `json:"lookup_bind_password,omitempty"`

	// server insecure
	ServerInsecure bool `json:"server_insecure,omitempty"`

	// server start tls
	ServerStartTLS bool `json:"server_start_tls,omitempty"`

	// skip tls verification
	SkipTLSVerification bool `json:"skip_tls_verification,omitempty"`

	// url
	// Required: true
	URL *string `json:"url"`

	// user dn search base dn
	UserDnSearchBaseDn string `json:"user_dn_search_base_dn,omitempty"`

	// user dn search filter
	UserDnSearchFilter string `json:"user_dn_search_filter,omitempty"`

	// user dns
	UserDNS []string `json:"user_dns"`

	// username format
	UsernameFormat string `json:"username_format,omitempty"`

	// username search filter
	UsernameSearchFilter string `json:"username_search_filter,omitempty"`
}

// Validate validates this idp configuration active directory
func (m *IdpConfigurationActiveDirectory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfigurationActiveDirectory) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("active_directory"+"."+"url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this idp configuration active directory based on context it is used
func (m *IdpConfigurationActiveDirectory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfigurationActiveDirectory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfigurationActiveDirectory) UnmarshalBinary(b []byte) error {
	var res IdpConfigurationActiveDirectory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdpConfigurationKeysItems0 idp configuration keys items0
//
// swagger:model IdpConfigurationKeysItems0
type IdpConfigurationKeysItems0 struct {

	// access key
	// Required: true
	AccessKey *string `json:"access_key"`

	// secret key
	// Required: true
	SecretKey *string `json:"secret_key"`
}

// Validate validates this idp configuration keys items0
func (m *IdpConfigurationKeysItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfigurationKeysItems0) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("access_key", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationKeysItems0) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secret_key", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this idp configuration keys items0 based on context it is used
func (m *IdpConfigurationKeysItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfigurationKeysItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfigurationKeysItems0) UnmarshalBinary(b []byte) error {
	var res IdpConfigurationKeysItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IdpConfigurationOidc idp configuration oidc
//
// swagger:model IdpConfigurationOidc
type IdpConfigurationOidc struct {

	// callback url
	CallbackURL string `json:"callback_url,omitempty"`

	// claim name
	// Required: true
	ClaimName *string `json:"claim_name"`

	// client id
	// Required: true
	ClientID *string `json:"client_id"`

	// configuration url
	// Required: true
	ConfigurationURL *string `json:"configuration_url"`

	// scopes
	Scopes string `json:"scopes,omitempty"`

	// secret id
	// Required: true
	SecretID *string `json:"secret_id"`
}

// Validate validates this idp configuration oidc
func (m *IdpConfigurationOidc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaimName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdpConfigurationOidc) validateClaimName(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"claim_name", "body", m.ClaimName); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationOidc) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationOidc) validateConfigurationURL(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"configuration_url", "body", m.ConfigurationURL); err != nil {
		return err
	}

	return nil
}

func (m *IdpConfigurationOidc) validateSecretID(formats strfmt.Registry) error {

	if err := validate.Required("oidc"+"."+"secret_id", "body", m.SecretID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this idp configuration oidc based on context it is used
func (m *IdpConfigurationOidc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdpConfigurationOidc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdpConfigurationOidc) UnmarshalBinary(b []byte) error {
	var res IdpConfigurationOidc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
