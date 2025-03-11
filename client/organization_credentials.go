package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	organizationCredentialsAPIEndpoint            = "/api/v2/organizations/%d/credentials/"
	organizationCredentialInputSourcesAPIEndpoint = "/api/v2/organizations/%d/credentials/%d/input_sources/"
)

// OrganizationCredentialResponse represents the credentials list response
type OrganizationCredentialResponse = PaginatedResponse[Credential]

// ListOrganizationCredentials shows list of credentials in an organization.
func (o *OrganizationsService) ListOrganizationCredentials(id int, params map[string]string) ([]*Credential, error) {
	result := new(OrganizationCredentialResponse)
	endpoint := fmt.Sprintf(organizationCredentialsAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationCredential creates a credential in the specified organization.
func (o *OrganizationsService) CreateOrganizationCredential(id int, data map[string]interface{}, params map[string]string) (*Credential, error) {
	mandatoryFields := []string{"name", "credential_type"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	// Ensure the organization field is set to the correct ID
	data["organization"] = id

	result := new(Credential)
	endpoint := fmt.Sprintf(organizationCredentialsAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateCredentialWithOrganization associates an existing credential with an organization
func (o *OrganizationsService) AssociateCredentialWithOrganization(organizationID int, credentialID int) error {
	data := map[string]interface{}{
		"id": credentialID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationCredentialsAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}

// DisassociateCredentialFromOrganization removes a credential's association with an organization
func (o *OrganizationsService) DisassociateCredentialFromOrganization(organizationID int, credentialID int) error {
	data := map[string]interface{}{
		"id":           credentialID,
		"disassociate": true,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationCredentialsAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}

// GetOrganizationCredentialInputSources retrieves the input sources for a credential
func (o *OrganizationsService) GetOrganizationCredentialInputSources(organizationID int, credentialID int, params map[string]string) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf(organizationCredentialInputSourcesAPIEndpoint, organizationID, credentialID)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
