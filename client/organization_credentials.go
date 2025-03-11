package awx

const organizationCredentials = "credentials"

// OrganizationCredentialResponse represents the credentials list response
type OrganizationCredentialResponse = PaginatedResponse[Credential]

// ListOrganizationCredentials shows list of credentials in an organization.
func (o *OrganizationsService) ListOrganizationCredentials(id int, params map[string]string) ([]*Credential, error) {
	return listOrganizationResource[Credential](o, id, organizationCredentials, params)
}

// CreateOrganizationCredential creates a credential in the specified organization.
func (o *OrganizationsService) CreateOrganizationCredential(id int, data map[string]interface{}) (*Credential, error) {
	mandatoryFields := []string{"name", "credential_type"}
	return createOrganizationResource[Credential](o, id, organizationCredentials, data, mandatoryFields)
}

// AssociateCredentialWithOrganization associates an existing credential with an organization
func (o *OrganizationsService) AssociateCredentialWithOrganization(organizationID int, credentialID int) error {
	data := map[string]interface{}{
		"id": credentialID,
	}

	_, err := o.associate(organizationID, organizationCredentials, data, nil)
	return err
}

// DisassociateCredentialFromOrganization removes a credential's association with an organization
func (o *OrganizationsService) DisassociateCredentialFromOrganization(organizationID int, credentialID int) error {
	data := map[string]interface{}{
		"id": credentialID,
	}

	_, err := o.disAssociate(organizationID, organizationCredentials, data, nil)
	return err
}
