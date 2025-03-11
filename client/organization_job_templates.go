package awx

const organizationJobTemplates = "job_templates"

// ListOrganizationJobTemplates shows list of job templates in an organization.
func (o *OrganizationsService) ListOrganizationJobTemplates(id int, params map[string]string) ([]*JobTemplate, error) {
	return listOrganizationResource[JobTemplate](o, id, organizationJobTemplates, params)
}

// CreateOrganizationJobTemplate creates a job template in the specified organization.
func (o *OrganizationsService) CreateOrganizationJobTemplate(id int, data map[string]interface{}) (*JobTemplate, error) {
	mandatoryFields := []string{"name", "job_type", "inventory", "project"}
	return createOrganizationResource[JobTemplate](o, id, organizationJobTemplates, data, mandatoryFields)
}

// AssociateJobTemplateWithOrganization associates an existing job template with an organization
func (o *OrganizationsService) AssociateJobTemplateWithOrganization(organizationID int, templateID int) error {
	data := map[string]interface{}{
		"id": templateID,
	}

	_, err := o.associate(organizationID, organizationJobTemplates, data, nil)
	return err
}

// DisassociateJobTemplateFromOrganization removes a job template's association with an organization
func (o *OrganizationsService) DisassociateJobTemplateFromOrganization(organizationID int, templateID int) error {
	data := map[string]interface{}{
		"id": templateID,
	}

	_, err := o.disAssociate(organizationID, organizationJobTemplates, data, nil)
	return err
}
