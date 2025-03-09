package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationUserResponse represents the users list response
type OrganizationUserResponse struct {
    Pagination
    Results []*User `json:"results"` // Using the base User type since the response matches
}

// ListOrganizationUsers shows list of users in an organization.
func (p *OrganizationsService) ListOrganizationUsers(id int, params map[string]string) ([]*User, error) {
    result := new(OrganizationUserResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/users/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// CreateOrganizationUser creates a user in the specified organization.
func (p *OrganizationsService) CreateOrganizationUser(id int, data map[string]interface{}, params map[string]string) (*User, error) {
    mandatoryFields := []string{"username", "password"}
    validate, status := ValidateParams(data, mandatoryFields)
    if !status {
        err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    result := new(User)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/users/", id)
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

// GetOrganizationUser retrieves a specific user in the organization
func (p *OrganizationsService) GetOrganizationUser(organizationID int, userID int, params map[string]string) (*User, error) {
    result := new(User)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/users/%d/", organizationID, userID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateOrganizationUser updates a user in the organization
func (p *OrganizationsService) UpdateOrganizationUser(organizationID int, userID int, data map[string]interface{}, params map[string]string) (*User, error) {
    result := new(User)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/users/%d/", organizationID, userID)
    
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

// DeleteOrganizationUser removes a user from the organization
func (p *OrganizationsService) DeleteOrganizationUser(organizationID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/users/%d/", organizationID, userID)
    
    resp, err := p.client.Requester.Delete(endpoint, nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}
