// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Swaggeroauth2TokenResponse The Access Token Response
// swagger:model swaggeroauth2TokenResponse
type Swaggeroauth2TokenResponse struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// Id token
	IDToken string `json:"id_token,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this swaggeroauth2 token response
func (m *Swaggeroauth2TokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Swaggeroauth2TokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Swaggeroauth2TokenResponse) UnmarshalBinary(b []byte) error {
	var res Swaggeroauth2TokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
