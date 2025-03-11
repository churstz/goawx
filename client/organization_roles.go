package awx

const organizationObjectRoles = "object_roles"

// ListOrganizationObjectRoles shows list of all object roles for an organization
func (o *OrganizationsService) ListOrganizationObjectRoles(id int, params map[string]string) ([]*Role, error) {
	return listOrganizationResource[Role](o, id, organizationObjectRoles, params)
}

// CreateOrganizationObjectRole creates an objectRole in the specified organization.
func (o *OrganizationsService) CreateOrganizationObjectRole(id int, data map[string]interface{}) (*Role, error) {
	mandatoryFields := []string{"name"}
	return createOrganizationResource[Role](o, id, organizationObjectRoles, data, mandatoryFields)
}

// AssociateRoleWithOrganization associates an existing role with an organization
func (o *OrganizationsService) AssociateRoleWithOrganization(organizationID int, roleID int) error {
	data := map[string]interface{}{
		"id": roleID,
	}

	_, err := o.associate(organizationID, organizationObjectRoles, data, nil)
	return err
}

// DisassociateRoleFromOrganization removes a role's association with an organization
func (o *OrganizationsService) DisassociateRoleFromOrganization(organizationID int, roleID int) error {
	data := map[string]interface{}{
		"id": roleID,
	}

	_, err := o.disAssociate(organizationID, organizationObjectRoles, data, nil)
	return err
}
