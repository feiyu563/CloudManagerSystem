package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ChartVersion chart version

swagger:model chartVersion
*/
type ChartVersion struct {

	/* app version

	Required: true
	Min Length: 1
	*/
	AppVersion *string `json:"app_version"`

	/* created

	Required: true
	Min Length: 1
	*/
	Created *string `json:"created"`

	/* digest

	Required: true
	Min Length: 1
	*/
	Digest *string `json:"digest"`

	/* icons
	 */
	Icons []*Icon `json:"icons,omitempty"`

	/* readme

	Required: true
	Min Length: 1
	*/
	Readme *string `json:"readme"`

	/* urls

	Required: true
	*/
	Urls []string `json:"urls"`

	/* version

	Required: true
	Min Length: 1
	*/
	Version *string `json:"version"`
}

// Validate validates this chart version
func (m *ChartVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDigest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIcons(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReadme(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUrls(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartVersion) validateAppVersion(formats strfmt.Registry) error {

	if err := validate.Required("app_version", "body", m.AppVersion); err != nil {
		return err
	}

	if err := validate.MinLength("app_version", "body", string(*m.AppVersion), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.MinLength("created", "body", string(*m.Created), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateDigest(formats strfmt.Registry) error {

	if err := validate.Required("digest", "body", m.Digest); err != nil {
		return err
	}

	if err := validate.MinLength("digest", "body", string(*m.Digest), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateIcons(formats strfmt.Registry) error {

	if swag.IsZero(m.Icons) { // not required
		return nil
	}

	for i := 0; i < len(m.Icons); i++ {

		if swag.IsZero(m.Icons[i]) { // not required
			continue
		}

		if m.Icons[i] != nil {

			if err := m.Icons[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ChartVersion) validateReadme(formats strfmt.Registry) error {

	if err := validate.Required("readme", "body", m.Readme); err != nil {
		return err
	}

	if err := validate.MinLength("readme", "body", string(*m.Readme), 1); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateUrls(formats strfmt.Registry) error {

	if err := validate.Required("urls", "body", m.Urls); err != nil {
		return err
	}

	return nil
}

func (m *ChartVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}
