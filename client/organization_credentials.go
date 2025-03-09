package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationCredentialResponse represents the credentials list response
type OrganizationCredentialResponse struct {
    Pagination
    Results []*Credential `json:"results"` // Using the base Credential type since the response matches
}

// ListOrganizationCredentials shows list of credentials in an organization.
func (p *OrganizationsService) ListOrganizationCredentials(id int, params map[string]string) ([]*Credential, error) {
    result := new(OrganizationCredentialResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/credentials/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// CreateOrganizationCredential creates a credential in the specified organization.
func (p *OrganizationsService) CreateOrganizationCredential(id int, data map[string]interface{}, params map[string]string) (*Credential, error) {
    mandatoryFields := []string{"name", "credential_type"}
    validate, status := ValidateParams(data, mandatoryFields)
    if !status {
        err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    // Ensure the organization field is set to the correct ID
    data["organization"] = id

    result := new(Credential)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/credentials/", id)
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// GetOrganizationCredential retrieves a specific credential in the organization
func (p *OrganizationsService) GetOrganizationCredential(organizationID int, credentialID int, params map[string]string) (*Credential, error) {
    result := new(Credential)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/credentials/%d/", organizationID, credentialID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateOrganizationCredential updates a credential in the organization
func (p *OrganizationsService) UpdateOrganizationCredential(organizationID int, credentialID int, data map[string]interface{}, params map[string]string) (*Credential, error) {
    result := new(Credential)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/credentials/%d/", organizationID, credentialID)
    
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// DeleteOrganizationCredential removes a credential from the organization
func (p *OrganizationsService) DeleteOrganizationCredential(organizationID int, credentialID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/credentials/%d/", organizationID, credentialID)
    
    resp, err := p.client.Requester.Delete(endpoint, nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// GetOrganizationCredentialInputSources retrieves the input sources for a credential
func (p *OrganizationsService) GetOrganizationCredentialInputSources(organizationID int, credentialID int, params map[string]string) (*Credential, error) {
    result := new(Credential)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/credentials/%d/input_sources/", organizationID, credentialID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}
