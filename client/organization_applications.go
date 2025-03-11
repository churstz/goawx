package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationApplicationsAPIEndpoint = "/api/v2/organizations/%d/applications/"

// OrganizationApplicationResponse represents the OAuth2 applications list response
type OrganizationApplicationResponse = PaginatedResponse[Application]

// ListOrganizationApplications shows list of OAuth2 applications in an organization.
func (o *OrganizationsService) ListOrganizationApplications(id int, params map[string]string) ([]*Application, error) {
	result := new(OrganizationApplicationResponse)
	endpoint := fmt.Sprintf(organizationApplicationsAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationApplication creates an OAuth2 application in the specified organization.
func (o *OrganizationsService) CreateOrganizationApplication(id int, data map[string]interface{}, params map[string]string) (*Application, error) {
	mandatoryFields := []string{"name", "client_type", "authorization_grant_type"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	// Ensure the organization field is set to the correct ID
	data["organization"] = id

	result := new(Application)
	endpoint := fmt.Sprintf(organizationApplicationsAPIEndpoint, id)
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
