package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationObjectRolesAPIEndpoint = "/api/v2/organizations/%d/object_roles/"

// OrganizationRoleResponse represents the organization roles list response
type OrganizationRoleResponse = PaginatedResponse[Role]

// ListOrganizationObjectRoles shows list of all object roles for an organization
func (o *OrganizationsService) ListOrganizationObjectRoles(id int, params map[string]string) ([]*Role, error) {
	result := new(OrganizationRoleResponse)
	endpoint := fmt.Sprintf(organizationObjectRolesAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationObjectRole creates an objectRole in the specified organization.
func (o *OrganizationsService) CreateOrganizationObjectRole(id int, data map[string]interface{}, params map[string]string) (*Role, error) {
	mandatoryFields := []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	// Ensure the organization field is set to the correct ID
	data["organization"] = id

	result := new(Role)
	endpoint := fmt.Sprintf(organizationObjectRolesAPIEndpoint, id)
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

// AssociateRoleWithOrganization associates an existing role with an organization
func (o *OrganizationsService) AssociateRoleWithOrganization(organizationID int, roleID int) error {
	data := map[string]interface{}{
		"id": roleID,
	}

	_, err := o.associate(organizationID, "roles", data, nil)
	return err
}

// DisassociateRoleFromOrganization removes a role's association with an organization
func (o *OrganizationsService) DisassociateRoleFromOrganization(organizationID int, roleID int) error {
	data := map[string]interface{}{
		"id": roleID,
	}

	_, err := o.disAssociate(organizationID, "roles", data, nil)
	return err
}
