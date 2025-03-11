package awx

const organizationUsers = "users"

// ListOrganizationUsers shows list of users in an organization.
func (o *OrganizationsService) ListOrganizationUsers(id int, params map[string]string) ([]*User, error) {
	return listOrganizationResource[User](o, id, organizationUsers, params)
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
func (o *OrganizationsService) CreateOrganizationUser(id int, data map[string]interface{}) (*User, error) {
	mandatoryFields := []string{"username", "email"}
	return createOrganizationResource[User](o, id, organizationUsers, data, mandatoryFields)
}

// AssociateUserWithOrganization associates an existing user with an organization
func (o *OrganizationsService) AssociateUserWithOrganization(organizationID int, userID int) error {
	data := map[string]interface{}{
		"id": userID,
	}

	_, err := o.associate(organizationID, organizationUsers, data, nil)
	return err
}

// DisassociateUserFromOrganization removes a user's association with an organization
func (o *OrganizationsService) DisassociateUserFromOrganization(organizationID int, userID int) error {
	data := map[string]interface{}{
		"id": userID,
	}

	_, err := o.disAssociate(organizationID, organizationUsers, data, nil)
	return err
}
