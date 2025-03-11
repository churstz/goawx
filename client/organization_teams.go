package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationTeamsAPIEndpoint = "/api/v2/organizations/%d/teams/"

// OrganizationTeamResponse represents the teams list response
type OrganizationTeamResponse = PaginatedResponse[Team]

// ListOrganizationTeams shows list of teams in an organization.
func (o *OrganizationsService) ListOrganizationTeams(id int, params map[string]string) ([]*Team, error) {
	result := new(OrganizationTeamResponse)
	endpoint := fmt.Sprintf(organizationTeamsAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationTeam creates a new team in the specified organization.
// Supported fields in data are:
// * name: (required) Name of the team.
// * description: (optional) Optional description of the team.
func (o *OrganizationsService) CreateOrganizationTeam(id int, data map[string]interface{}, params map[string]string) (*Team, error) {
	mandatoryFields := []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	// Ensure the organization field is set to the correct ID
	data["organization"] = id

	result := new(Team)
	endpoint := fmt.Sprintf(organizationTeamsAPIEndpoint, id)
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

// AssociateTeamWithOrganization associates an existing team with an organization
func (o *OrganizationsService) AssociateTeamWithOrganization(organizationID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := o.associate(organizationID, "teams", data, nil)
	return err
}

// DisassociateTeamFromOrganization removes a team's association with an organization
func (o *OrganizationsService) DisassociateTeamFromOrganization(organizationID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := o.disAssociate(organizationID, "teams", data, nil)
	return err
}
