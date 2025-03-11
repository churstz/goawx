package awx

const organizationAdmins = "admins"

func (o *OrganizationsService) ListOrganizationAdmins(id int, params map[string]string) ([]*User, error) {
	return listOrganizationResource[User](o, id, organizationAdmins, params)
}

func (o *OrganizationsService) CreateOrganizationAdmin(orgID int, data map[string]interface{}) (*User, error) {
	mandatoryFields := []string{"username", "email"}
	return createOrganizationResource[User](o, orgID, organizationAdmins, data, mandatoryFields)
}

// AssociateOrganizationAdmin makes a user an admin of the organization
func (o *OrganizationsService) AssociateOrganizationAdmin(organizationID int, adminID int) error {
	data := map[string]interface{}{
		"id": adminID,
	}

	_, err := o.associate(organizationID, organizationAdmins, data, nil)
	return err
}

// DisassociateOrganizationAdmin removes a user's admin role from the organization
func (o *OrganizationsService) DisassociateOrganizationAdmin(organizationID int, adminID int) error {
	data := map[string]interface{}{
		"id": adminID,
	}

	_, err := o.disAssociate(organizationID, organizationAdmins, data, nil)
	return err
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
