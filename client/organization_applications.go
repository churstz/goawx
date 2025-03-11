package awx

const organizationApplications = "applications"

// OrganizationApplicationResponse represents the OAuth2 applications list response
type OrganizationApplicationResponse = PaginatedResponse[Application]

// ListOrganizationApplications shows list of OAuth2 applications in an organization.
func (o *OrganizationsService) ListOrganizationApplications(id int, params map[string]string) ([]*Application, error) {
	return listOrganizationResource[Application](o, id, organizationApplications, params)
}

// CreateOrganizationApplication creates an OAuth2 application in the specified organization.
func (o *OrganizationsService) CreateOrganizationApplication(id int, data map[string]interface{}) (*Application, error) {
	mandatoryFields := []string{"name", "client_type", "authorization_grant_type"}
	return createOrganizationResource[Application](o, id, organizationApplications, data, mandatoryFields)
}

// AssociateApplicationWithOrganization associates an existing application with an organization
func (o *OrganizationsService) AssociateApplicationWithOrganization(organizationID int, applicationID int) error {
	data := map[string]interface{}{
		"id": applicationID,
	}

	_, err := o.associate(organizationID, "applications", data, nil)
	return err
}

// DisassociateApplicationFromOrganization removes an application's association with an organization
func (o *OrganizationsService) DisassociateApplicationFromOrganization(organizationID int, applicationID int) error {
	data := map[string]interface{}{
		"id": applicationID,
	}

	_, err := o.disAssociate(organizationID, "applications", data, nil)
	return err
}
