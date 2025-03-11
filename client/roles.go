package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	// Resource Types
	ResourceTypeProject     = "projects"
	ResourceTypeInventory   = "inventories"
	ResourceTypeJobTemplate = "job_templates"
	ResourceTypeWorkflow    = "workflow_job_templates"
	ResourceTypeCredential  = "credentials"
)

// RoleService implements role management operations
type RoleService struct {
	client *Client
}

// RoleSummaryFields represents summary fields in role response
type RoleSummaryFields struct {
	ResourceName            string `json:"resource_name"`
	ResourceType            string `json:"resource_type"`
	ResourceTypeDisplayName string `json:"resource_type_display_name"`
	ResourceID              int    `json:"resource_id"`
}

// Role represents an AWX role
type Role struct {
	ID            int               `json:"id"`
	Type          string            `json:"type"`
	URL           string            `json:"url"`
	Related       map[string]string `json:"related"`
	SummaryFields RoleSummaryFields `json:"summary_fields"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
}

// RoleListResponse represents the roles list response
type RoleListResponse = PaginatedResponse[Role]

const rolesAPIEndpoint = "/api/v2/roles/"

// ListRoles retrieves a list of all roles with pagination
func (r *RoleService) ListRoles(params map[string]string) ([]*Role, error) {
	results, err := r.getAllPages(rolesAPIEndpoint, params)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetRole retrieves a specific role by ID
func (r *RoleService) GetRoleByID(id int, params map[string]string) (*Role, error) {
	result := new(Role)
	endpoint := fmt.Sprintf("/api/v2/roles/%d/", id)

	resp, err := r.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateRole creates an awx Role
func (p *RoleService) CreateRole(data map[string]interface{}, params map[string]string) (*Role, error) {
	mandatoryFields = []string{"name"} //TODO: Check this
	validate, status := ValidateParams(data, mandatoryFields)

	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(Role)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(rolesAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateRole update an awx Role.
func (p *RoleService) UpdateRole(id int, data map[string]interface{}, params map[string]string) (*Role, error) {
	result := new(Role)
	endpoint := fmt.Sprintf("%s%d", rolesAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteRole delete an awx Organization.
func (p *RoleService) DeleteRole(id int) (*Role, error) {
	result := new(Role)
	endpoint := fmt.Sprintf("%s%d", rolesAPIEndpoint, id)

	resp, err := p.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// Organization "Related" Resources.  For us when adding or removing an existing resource to/from an organization using POST and an ID
// Associate associate an element to Role.
func (p *RoleService) associate(id int, typ string, data map[string]interface{}, params map[string]string) (*Role, error) {
	result := new(Role)

	endpoint := fmt.Sprintf("%s%d/%s/", rolesAPIEndpoint, id, typ)
	data["associate"] = true
	mandatoryFields = []string{"id"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DisAssociate remove element from a role
func (p *RoleService) disAssociate(id int, typ string, data map[string]interface{}, params map[string]string) (*Role, error) {
	result := new(Role)
	endpoint := fmt.Sprintf("%s%d/%s/", rolesAPIEndpoint, id, typ)
	data["disassociate"] = true
	mandatoryFields = []string{"id", "disassociate"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *RoleService) getAllPages(firstURL string, params map[string]string) ([]*Role, error) {
	results := make([]*Role, 0)
	nextURL := firstURL
	for {
		nextURLParsed, err := url.Parse(nextURL)
		if err != nil {
			return nil, err
		}

		nextURLQueryParams := make(map[string]string)
		for paramName, paramValues := range nextURLParsed.Query() {
			if len(paramValues) > 0 {
				nextURLQueryParams[paramName] = paramValues[0]
			}
		}

		for paramName, paramValue := range params {
			nextURLQueryParams[paramName] = paramValue
		}

		result := new(PaginatedResponse[Role])
		resp, err := p.client.Requester.GetJSON(nextURLParsed.Path, result, nextURLQueryParams)
		if err != nil {
			return nil, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, err
		}

		results = append(results, result.Results...)

		if result.Next == nil || result.Next.(string) == "" {
			break
		}
		nextURL = result.Next.(string)
	}
	return results, nil
}

// #########################################
const roleTeams = "teams"

func (r *RoleService) ListRoleTeams(id int, params map[string]string) ([]*Team, error) {
	return listRoleResource[Team](r, id, roleTeams, params)
}

func (r *RoleService) CreateRoleTeam(id int, data map[string]interface{}) (*Team, error) {
	mandatoryFields := []string{"name"}
	return createRoleResource[Team](r, id, roleTeams, data, mandatoryFields)
}

func (r *RoleService) AssociateTeamWithRole(roleID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := r.associate(roleID, roleTeams, data, nil)
	return err
}

func (r *RoleService) DisassociateTeamFromRole(roleID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := r.disAssociate(roleID, roleTeams, data, nil)
	return err
}

// #########################################
const roleUsers = "users"

func (r *RoleService) ListRoleUsers(id int, params map[string]string) ([]*User, error) {
	return listRoleResource[User](r, id, roleUsers, params)
}

func (r *RoleService) CreateRoleUser(id int, data map[string]interface{}) (*User, error) {
	mandatoryFields := []string{"name"}
	return createRoleResource[User](r, id, roleUsers, data, mandatoryFields)
}

func (r *RoleService) AssociateUserithRole(roleID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := r.associate(roleID, roleUsers, data, nil)
	return err
}

func (r *RoleService) DisassociateUserFromRole(roleID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := r.disAssociate(roleID, roleUsers, data, nil)
	return err
}

// #########################################
const roleJobTemplate = "job_templates"

func (r *RoleService) ListRoleJobTemplates(id int, params map[string]string) ([]*JobTemplate, error) {
	return listRoleResource[JobTemplate](r, id, roleJobTemplate, params)
}

func (r *RoleService) CreateRoleJobTemplate(id int, data map[string]interface{}) (*JobTemplate, error) {
	mandatoryFields := []string{"name"}
	return createRoleResource[JobTemplate](r, id, roleJobTemplate, data, mandatoryFields)
}

func (r *RoleService) AssociateJobTemplateWithRole(jobTemplateID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := r.associate(jobTemplateID, roleJobTemplate, data, nil)
	return err
}

func (r *RoleService) DisassociateJobTemplateFromRole(jobTemplateID int, teamID int) error {
	data := map[string]interface{}{
		"id": teamID,
	}

	_, err := r.disAssociate(jobTemplateID, roleJobTemplate, data, nil)
	return err
}

//#########################################
//Free Functions for Role Resources
// (not part of the RoleService struct)
//#########################################

func getRolesAllResourcePages[T any](p *RoleService, firstURL string, params map[string]string) ([]*T, error) {
	results := make([]*T, 0)
	nextURL := firstURL
	for {
		nextURLParsed, err := url.Parse(nextURL)
		if err != nil {
			return nil, err
		}

		nextURLQueryParams := make(map[string]string)
		for paramName, paramValues := range nextURLParsed.Query() {
			if len(paramValues) > 0 {
				nextURLQueryParams[paramName] = paramValues[0]
			}
		}

		for paramName, paramValue := range params {
			nextURLQueryParams[paramName] = paramValue
		}

		result := new(PaginatedResponse[T])
		resp, err := p.client.Requester.GetJSON(nextURLParsed.Path, result, nextURLQueryParams)
		if err != nil {
			return nil, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, err
		}

		results = append(results, result.Results...)

		if result.Next == nil || result.Next.(string) == "" {
			break
		}
		nextURL = result.Next.(string)
	}
	return results, nil
}

func listRoleResource[T any](p *RoleService, roleID int, resourceType string, params map[string]string) ([]*T, error) {
	endpoint := fmt.Sprintf("%s%d/%s/", rolesAPIEndpoint, roleID, resourceType)

	results, err := getRolesAllResourcePages[T](p, endpoint, params)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func createRoleResource[T any](p *RoleService, roleID int, resourceType string, data map[string]interface{}, mandatoryFields []string) (*T, error) {
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	result := new(T)
	endpoint := fmt.Sprintf("%s%d/%s/", rolesAPIEndpoint, roleID, resourceType)

	// Add role ID to data if not present
	if _, ok := data["role"]; !ok {
		data["role"] = roleID
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
