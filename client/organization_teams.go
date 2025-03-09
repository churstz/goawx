package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationTeamResponse represents the teams list response
type OrganizationTeamResponse struct {
    Pagination
    Results []*Team `json:"results"` // Using the base Team type since the response matches
}

// ListOrganizationTeams shows list of teams in an organization.
func (p *OrganizationsService) ListOrganizationTeams(id int, params map[string]string) ([]*Team, error) {
    result := new(OrganizationTeamResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/teams/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// CreateOrganizationTeam creates a team in the specified organization.
func (p *OrganizationsService) CreateOrganizationTeam(id int, data map[string]interface{}, params map[string]string) (*Team, error) {
    mandatoryFields := []string{"name"}
    validate, status := ValidateParams(data, mandatoryFields)
    if !status {
        err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    // Ensure the organization field is set to the correct ID
    data["organization"] = id

    result := new(Team)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/teams/", id)
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

// GetOrganizationTeam retrieves a specific team in the organization
func (p *OrganizationsService) GetOrganizationTeam(organizationID int, teamID int, params map[string]string) (*Team, error) {
    result := new(Team)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/teams/%d/", organizationID, teamID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateOrganizationTeam updates a team in the organization
func (p *OrganizationsService) UpdateOrganizationTeam(organizationID int, teamID int, data map[string]interface{}, params map[string]string) (*Team, error) {
    result := new(Team)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/teams/%d/", organizationID, teamID)
    
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

// DeleteOrganizationTeam removes a team from the organization
func (p *OrganizationsService) DeleteOrganizationTeam(organizationID int, teamID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/teams/%d/", organizationID, teamID)
    
    resp, err := p.client.Requester.Delete(endpoint, nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}
