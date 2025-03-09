package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationJobTemplateResponse represents the job templates list response
type OrganizationJobTemplateResponse struct {
    Pagination
    Results []*JobTemplate `json:"results"` // Using the base JobTemplate type since the response matches
}

// ListOrganizationJobTemplates shows list of job templates in an organization.
func (p *OrganizationsService) ListOrganizationJobTemplates(id int, params map[string]string) ([]*JobTemplate, error) {
    result := new(OrganizationJobTemplateResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/job_templates/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// CreateOrganizationJobTemplate creates a job template in the specified organization.
func (p *OrganizationsService) CreateOrganizationJobTemplate(id int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
    mandatoryFields := []string{"name", "job_type", "inventory", "project"}
    validate, status := ValidateParams(data, mandatoryFields)
    if !status {
        err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    // Ensure the organization field is set to the correct ID
    data["organization"] = id

    result := new(JobTemplate)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/job_templates/", id)
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// GetOrganizationJobTemplate retrieves a specific job template in the organization
func (p *OrganizationsService) GetOrganizationJobTemplate(organizationID int, templateID int, params map[string]string) (*JobTemplate, error) {
    result := new(JobTemplate)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/job_templates/%d/", organizationID, templateID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateOrganizationJobTemplate updates a job template in the organization
func (p *OrganizationsService) UpdateOrganizationJobTemplate(organizationID int, templateID int, data map[string]interface{}, params map[string]string) (*JobTemplate, error) {
    result := new(JobTemplate)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/job_templates/%d/", organizationID, templateID)
    
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// DeleteOrganizationJobTemplate removes a job template from the organization
func (p *OrganizationsService) DeleteOrganizationJobTemplate(organizationID int, templateID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/job_templates/%d/", organizationID, templateID)
    
    resp, err := p.client.Requester.Delete(endpoint, nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// LaunchOrganizationJobTemplate launches a job with the organization's job template
func (p *OrganizationsService) LaunchOrganizationJobTemplate(organizationID int, templateID int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
    result := new(JobLaunch)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/job_templates/%d/launch/", organizationID, templateID)
    
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}
