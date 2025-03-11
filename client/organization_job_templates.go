package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationJobTemplatesAPIEndpoint = "/api/v2/organizations/%d/job_templates/"

// OrganizationJobTemplateResponse represents the job templates list response
type OrganizationJobTemplateResponse = PaginatedResponse[JobTemplate]

// ListOrganizationJobTemplates shows list of job templates in an organization.
func (o *OrganizationsService) ListOrganizationJobTemplates(id int, params map[string]string) ([]*JobTemplate, error) {
	result := new(OrganizationJobTemplateResponse)
	endpoint := fmt.Sprintf(organizationJobTemplatesAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationJobTemplate creates a job template in the specified organization.
func (o *OrganizationsService) CreateOrganizationJobTemplate(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
	mandatoryFields := []string{"name", "job_type", "inventory", "project"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	// Ensure the organization field is set to the correct ID
	data["organization"] = id

	result := new(JobTemplate)
	endpoint := fmt.Sprintf(organizationJobTemplatesAPIEndpoint, id)
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

// AssociateJobTemplateWithOrganization associates an existing job template with an organization
func (o *OrganizationsService) AssociateJobTemplateWithOrganization(organizationID int, templateID int) error {
	data := map[string]interface{}{
		"id": templateID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationJobTemplatesAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}

// DisassociateJobTemplateFromOrganization removes a job template's association with an organization
func (o *OrganizationsService) DisassociateJobTemplateFromOrganization(organizationID int, templateID int) error {
	data := map[string]interface{}{
		"id":           templateID,
		"disassociate": true,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationJobTemplatesAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}
