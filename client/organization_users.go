package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationUsersAPIEndpoint = "/api/v2/organizations/%d/users/"

// OrganizationUserResponse represents the users list response
type OrganizationUserResponse = PaginatedResponse[User]

// ListOrganizationUsers shows list of users in an organization.
func (o *OrganizationsService) ListOrganizationUsers(id int, params map[string]string) ([]*User, error) {
	result := new(OrganizationUserResponse)
	endpoint := fmt.Sprintf(organizationUsersAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationUser creates a new user in the specified organization.
// Supported fields in data are:
// * username: (required) 150 characters or fewer. Letters, digits and './+/-/_' only.
// * password: (required) Field used to set the password.
// * first_name: (optional) First name of user.
// * last_name: (optional) Last name of user.
// * email: (optional) Email address.
// * is_superuser: (optional) Designates that this user has all permissions.
// * is_system_auditor: (optional) System auditor flag.
func (o *OrganizationsService) CreateOrganizationUser(id int, data map[string]interface{}, params map[string]string) (*User, error) {
	mandatoryFields := []string{"username", "password"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(User)
	endpoint := fmt.Sprintf(organizationUsersAPIEndpoint, id)
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

// AssociateUserWithOrganization associates an existing user with an organization
func (o *OrganizationsService) AssociateUserWithOrganization(organizationID int, userID int) error {
	data := map[string]interface{}{
		"id": userID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationUsersAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}

// DisassociateUserFromOrganization removes a user's association with an organization
func (o *OrganizationsService) DisassociateUserFromOrganization(organizationID int, userID int) error {
	data := map[string]interface{}{
		"id":           userID,
		"disassociate": true,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationUsersAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}
