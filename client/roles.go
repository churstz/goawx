package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

const (
    // Resource Types
    ResourceTypeProject      = "projects"
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
    ResourceName           string `json:"resource_name"`
    ResourceType          string `json:"resource_type"`
    ResourceTypeDisplayName string `json:"resource_type_display_name"`
    ResourceID           int    `json:"resource_id"`
}

// Role represents an AWX role
type Role struct {
    ID            int               `json:"id"`
    Type          string           `json:"type"`
    URL           string           `json:"url"`
    Related       map[string]string `json:"related"`
    SummaryFields RoleSummaryFields `json:"summary_fields"`
    Name          string           `json:"name"`
    Description   string           `json:"description"`
}

// RoleListResponse represents the roles list response
type RoleListResponse struct {
    Pagination
    Results []*Role `json:"results"`
}

// ListRoles retrieves a list of all roles with pagination
func (r *RoleService) ListRoles(params map[string]string) (*RoleListResponse, error) {
    result := new(RoleListResponse)
    endpoint := "/api/v2/roles/"
    
    resp, err := r.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// ListAllRoles retrieves all roles by handling pagination automatically
func (r *RoleService) ListAllRoles(params map[string]string) ([]*Role, error) {
    if params == nil {
        params = make(map[string]string)
    }
    
    var roles []*Role
    for {
        result, err := r.ListRoles(params)
        if err != nil {
            return nil, err
        }
        roles = append(roles, result.Results...)
        
        if result.Next == "" {
            break
        }
        params["page"] = result.NextPage
    }
    
    return roles, nil
}

// GetProjectRoles lists all roles for a project with pagination
func (r *RoleService) GetProjectRoles(projectID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceRoles(ResourceTypeProject, projectID, params)
}

// GetAllProjectRoles retrieves all roles for a project
func (r *RoleService) GetAllProjectRoles(projectID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceRoles(ResourceTypeProject, projectID, params)
}

// GetInventoryRoles lists all roles for an inventory with pagination
func (r *RoleService) GetInventoryRoles(inventoryID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceRoles(ResourceTypeInventory, inventoryID, params)
}

// GetAllInventoryRoles retrieves all roles for an inventory
func (r *RoleService) GetAllInventoryRoles(inventoryID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceRoles(ResourceTypeInventory, inventoryID, params)
}

// GetJobTemplateRoles lists all roles for a job template with pagination
func (r *RoleService) GetJobTemplateRoles(jobTemplateID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceRoles(ResourceTypeJobTemplate, jobTemplateID, params)
}

// GetAllJobTemplateRoles retrieves all roles for a job template
func (r *RoleService) GetAllJobTemplateRoles(jobTemplateID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceRoles(ResourceTypeJobTemplate, jobTemplateID, params)
}

// GetWorkflowRoles lists all roles for a workflow job template with pagination
func (r *RoleService) GetWorkflowRoles(workflowID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceRoles(ResourceTypeWorkflow, workflowID, params)
}

// GetAllWorkflowRoles retrieves all roles for a workflow job template
func (r *RoleService) GetAllWorkflowRoles(workflowID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceRoles(ResourceTypeWorkflow, workflowID, params)
}

// GetCredentialRoles lists all roles for a credential with pagination
func (r *RoleService) GetCredentialRoles(credentialID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceRoles(ResourceTypeCredential, credentialID, params)
}

// GetAllCredentialRoles retrieves all roles for a credential
func (r *RoleService) GetAllCredentialRoles(credentialID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceRoles(ResourceTypeCredential, credentialID, params)
}

// GetProjectObjectRoles lists all object roles for a project with pagination
func (r *RoleService) GetProjectObjectRoles(projectID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceObjectRoles(ResourceTypeProject, projectID, params)
}

// GetAllProjectObjectRoles retrieves all object roles for a project
func (r *RoleService) GetAllProjectObjectRoles(projectID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceObjectRoles(ResourceTypeProject, projectID, params)
}

// GetInventoryObjectRoles lists all object roles for an inventory with pagination
func (r *RoleService) GetInventoryObjectRoles(inventoryID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceObjectRoles(ResourceTypeInventory, inventoryID, params)
}

// GetAllInventoryObjectRoles retrieves all object roles for an inventory
func (r *RoleService) GetAllInventoryObjectRoles(inventoryID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceObjectRoles(ResourceTypeInventory, inventoryID, params)
}

// GetJobTemplateObjectRoles lists all object roles for a job template with pagination
func (r *RoleService) GetJobTemplateObjectRoles(jobTemplateID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceObjectRoles(ResourceTypeJobTemplate, jobTemplateID, params)
}

// GetAllJobTemplateObjectRoles retrieves all object roles for a job template
func (r *RoleService) GetAllJobTemplateObjectRoles(jobTemplateID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceObjectRoles(ResourceTypeJobTemplate, jobTemplateID, params)
}

// GetWorkflowObjectRoles lists all object roles for a workflow job template with pagination
func (r *RoleService) GetWorkflowObjectRoles(workflowID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceObjectRoles(ResourceTypeWorkflow, workflowID, params)
}

// GetAllWorkflowObjectRoles retrieves all object roles for a workflow job template
func (r *RoleService) GetAllWorkflowObjectRoles(workflowID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceObjectRoles(ResourceTypeWorkflow, workflowID, params)
}

// GetCredentialObjectRoles lists all object roles for a credential with pagination
func (r *RoleService) GetCredentialObjectRoles(credentialID int, params map[string]string) (*RoleListResponse, error) {
    return r.GetResourceObjectRoles(ResourceTypeCredential, credentialID, params)
}

// GetAllCredentialObjectRoles retrieves all object roles for a credential
func (r *RoleService) GetAllCredentialObjectRoles(credentialID int, params map[string]string) ([]*Role, error) {
    return r.GetAllResourceObjectRoles(ResourceTypeCredential, credentialID, params)
}

// GetResourceRoles lists all roles for a given resource with pagination
func (r *RoleService) GetResourceRoles(resourceType string, resourceID int, params map[string]string) (*RoleListResponse, error) {
    result := new(RoleListResponse)
    endpoint := fmt.Sprintf("/api/v2/%s/%d/roles/", resourceType, resourceID)
    
    resp, err := r.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// GetAllResourceRoles retrieves all roles for a given resource by handling pagination automatically
func (r *RoleService) GetAllResourceRoles(resourceType string, resourceID int, params map[string]string) ([]*Role, error) {
    if params == nil {
        params = make(map[string]string)
    }
    
    var roles []*Role
    for {
        result, err := r.GetResourceRoles(resourceType, resourceID, params)
        if err != nil {
            return nil, err
        }
        roles = append(roles, result.Results...)
        
        if result.Next == "" {
            break
        }
        params["page"] = result.NextPage
    }
    
    return roles, nil
}

// GetResourceObjectRoles lists all object roles for a given resource with pagination
func (r *RoleService) GetResourceObjectRoles(resourceType string, resourceID int, params map[string]string) (*RoleListResponse, error) {
    result := new(RoleListResponse)
    endpoint := fmt.Sprintf("/api/v2/%s/%d/object_roles/", resourceType, resourceID)
    
    resp, err := r.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// GetAllResourceObjectRoles retrieves all object roles for a given resource by handling pagination automatically
func (r *RoleService) GetAllResourceObjectRoles(resourceType string, resourceID int, params map[string]string) ([]*Role, error) {
    if params == nil {
        params = make(map[string]string)
    }
    
    var roles []*Role
    for {
        result, err := r.GetResourceObjectRoles(resourceType, resourceID, params)
        if err != nil {
            return nil, err
        }
        roles = append(roles, result.Results...)
        
        if result.Next == "" {
            break
        }
        params["page"] = result.NextPage
    }
    
    return roles, nil
}

// AssignUserRole assigns a user to a role
func (r *RoleService) AssignUserRole(roleID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/roles/%d/users/", roleID)
    data := map[string]interface{}{
        "id": userID,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := r.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// RemoveUserRole removes a user from a role
func (r *RoleService) RemoveUserRole(roleID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/roles/%d/users/", roleID)
    data := map[string]interface{}{
        "id": userID,
        "disassociate": true,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := r.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// AssignTeamRole assigns a team to a role
func (r *RoleService) AssignTeamRole(roleID int, teamID int) error {
    endpoint := fmt.Sprintf("/api/v2/roles/%d/teams/", roleID)
    data := map[string]interface{}{
        "id": teamID,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := r.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// RemoveTeamRole removes a team from a role
func (r *RoleService) RemoveTeamRole(roleID int, teamID int) error {
    endpoint := fmt.Sprintf("/api/v2/roles/%d/teams/", roleID)
    data := map[string]interface{}{
        "id": teamID,
        "disassociate": true,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := r.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// GetRole retrieves a specific role by ID
func (r *RoleService) GetRole(id int, params map[string]string) (*Role, error) {
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

// ListRoleUsers retrieves the list of users that have this role with pagination
func (r *RoleService) ListRoleUsers(id int, params map[string]string) (*UserListResponse, error) {
    result := new(UserListResponse)
    endpoint := fmt.Sprintf("/api/v2/roles/%d/users/", id)
    
    resp, err := r.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// ListAllRoleUsers retrieves all users that have this role by handling pagination automatically
func (r *RoleService) ListAllRoleUsers(id int, params map[string]string) ([]*User, error) {
    if params == nil {
        params = make(map[string]string)
    }
    
    var users []*User
    for {
        result, err := r.ListRoleUsers(id, params)
        if err != nil {
            return nil, err
        }
        users = append(users, result.Results...)
        
        if result.Next == "" {
            break
        }
        params["page"] = result.NextPage
    }
    
    return users, nil
}

// ListRoleTeams retrieves the list of teams that have this role with pagination
func (r *RoleService) ListRoleTeams(id int, params map[string]string) (*TeamListResponse, error) {
    result := new(TeamListResponse)
    endpoint := fmt.Sprintf("/api/v2/roles/%d/teams/", id)
    
    resp, err := r.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// ListAllRoleTeams retrieves all teams that have this role by handling pagination automatically
func (r *RoleService) ListAllRoleTeams(id int, params map[string]string) ([]*Team, error) {
    if params == nil {
        params = make(map[string]string)
    }
    
    var teams []*Team
    for {
        result, err := r.ListRoleTeams(id, params)
        if err != nil {
            return nil, err
        }
        teams = append(teams, result.Results...)
        
        if result.Next == "" {
            break
        }
        params["page"] = result.NextPage
    }
    
    return teams, nil
}
