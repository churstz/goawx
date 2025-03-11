package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationAdminsAPIEndpoint = "/api/v2/organizations/%d/admins/"

// OrganizationAdminResponse represents the admins list response
type OrganizationAdminResponse = PaginatedResponse[User]

// ListOrganizationAdmins shows list of admin users for an organization.
func (o *OrganizationsService) ListOrganizationAdmins(id int, params map[string]string) ([]*User, error) {
	result := new(OrganizationAdminResponse)
	endpoint := fmt.Sprintf(organizationAdminsAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// AssociateOrganizationAdmin makes a user an admin of the organization
func (o *OrganizationsService) AssociateOrganizationAdmin(organizationID int, userID int) ([]*User, error) {
	result := new(OrganizationAdminResponse)
	endpoint := fmt.Sprintf(organizationAdminsAPIEndpoint, organizationID)

	data := map[string]interface{}{
		"id":        userID,
		"associate": true,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// DisassociateOrganizationAdmin removes a user's admin role from the organization
func (o *OrganizationsService) DisassociateOrganizationAdmin(organizationID int, userID int) ([]*User, error) {
	result := new(OrganizationAdminResponse)
	endpoint := fmt.Sprintf(organizationAdminsAPIEndpoint, organizationID)

	data := map[string]interface{}{
		"id":           userID,
		"disassociate": true,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// IsOrganizationAdmin checks if a specific user is an admin of the organization
func (o *OrganizationsService) IsOrganizationAdmin(organizationID int, userID int) (bool, error) {
	admins, err := o.ListOrganizationAdmins(organizationID, nil)
	if err != nil {
		return false, err
	}

	for _, admin := range admins {
		if admin.ID == userID {
			return true, nil
		}
	}

	return false, nil
}
