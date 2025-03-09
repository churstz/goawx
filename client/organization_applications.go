package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationApplicationResponse represents the OAuth2 applications list response
type OrganizationApplicationResponse struct {
    Pagination
    Results []*Application `json:"results"` // Using the base Application type since the response matches
}

// ListOrganizationApplications shows list of OAuth2 applications in an organization.
func (p *OrganizationsService) ListOrganizationApplications(id int, params map[string]string) ([]*Application, error) {
    result := new(OrganizationApplicationResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/applications/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// CreateOrganizationApplication creates an OAuth2 application in the specified organization.
func (p *OrganizationsService) CreateOrganizationApplication(id int, data map[string]interface{}, params map[string]string) (*Application, error) {
    mandatoryFields := []string{"name", "client_type", "authorization_grant_type"}
    validate, status := ValidateParams(data, mandatoryFields)
    if !status {
        err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    // Ensure the organization field is set to the correct ID
    data["organization"] = id

    result := new(Application)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/applications/", id)
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

// GetOrganizationApplication retrieves a specific OAuth2 application in the organization
func (p *OrganizationsService) GetOrganizationApplication(organizationID int, applicationID int, params map[string]string) (*Application, error) {
    result := new(Application)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/applications/%d/", organizationID, applicationID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateOrganizationApplication updates an OAuth2 application in the organization
func (p *OrganizationsService) UpdateOrganizationApplication(organizationID int, applicationID int, data map[string]interface{}, params map[string]string) (*Application, error) {
    result := new(Application)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/applications/%d/", organizationID, applicationID)
    
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

// DeleteOrganizationApplication removes an OAuth2 application from the organization
func (p *OrganizationsService) DeleteOrganizationApplication(organizationID int, applicationID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/applications/%d/", organizationID, applicationID)
    
    resp, err := p.client.Requester.Delete(endpoint, nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}
