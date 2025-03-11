package awx

const organizationTeams = "teams"

// ListOrganizationTeams shows list of teams in an organization.
func (o *OrganizationsService) ListOrganizationTeams(id int, params map[string]string) ([]*Team, error) {
	return listOrganizationResource[Team](o, id, organizationTeams, params)
}

// CreateOrganizationTeam creates a new team in the specified organization.
// Supported fields in data are:
// * name: (required) Name of the team.
// * description: (optional) Optional description of the team.
func (o *OrganizationsService) CreateOrganizationTeam(id int, data map[string]interface{}) (*Team, error) {
	mandatoryFields := []string{"name"}
	return createOrganizationResource[Team](o, id, organizationTeams, data, mandatoryFields)
}

// AssociateTeamWithOrganization associates an existing team with an organization
func (o *OrganizationsService) AssociateTeamWithOrganization(organizationID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := o.associate(organizationID, organizationTeams, data, nil)
	return err
}

// DisassociateTeamFromOrganization removes a team's association with an organization
func (o *OrganizationsService) DisassociateTeamFromOrganization(organizationID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := o.disAssociate(organizationID, organizationTeams, data, nil)
	return err
}
