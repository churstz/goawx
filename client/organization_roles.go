package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// ListOrganizationRoles shows list of all roles for an organization
func (o *OrganizationsService) ListOrganizationRoles(id int, params map[string]string) ([]*Role, error) {
    result := new(RoleListResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/roles/", id)
    
    resp, err := o.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// ListOrganizationObjectRoles shows list of all object roles for an organization
func (o *OrganizationsService) ListOrganizationObjectRoles(id int, params map[string]string) ([]*Role, error) {
    result := new(RoleListResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/object_roles/", id)
    
    resp, err := o.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// GetOrganizationRole retrieves a specific role for an organization
func (o *OrganizationsService) GetOrganizationRole(organizationID int, roleID int, params map[string]string) (*Role, error) {
    return o.client.Roles.GetRole(roleID, params)
}

// ListOrganizationRoleUsers shows list of users that have been assigned a specific role for this organization
func (o *OrganizationsService) ListOrganizationRoleUsers(organizationID int, roleID int, params map[string]string) ([]*User, error) {
    return o.client.Roles.ListRoleUsers(roleID, params)
}

// ListOrganizationRoleTeams shows list of teams that have been assigned a specific role for this organization
func (o *OrganizationsService) ListOrganizationRoleTeams(organizationID int, roleID int, params map[string]string) ([]*Team, error) {
    return o.client.Roles.ListRoleTeams(roleID, params)
}

// AssignUserOrganizationRole assigns a user to a role within an organization
func (o *OrganizationsService) AssignUserOrganizationRole(organizationID int, roleID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/roles/%d/users/", organizationID, roleID)
    data := map[string]interface{}{
        "id": userID,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// RemoveUserOrganizationRole removes a user from a role within an organization
func (o *OrganizationsService) RemoveUserOrganizationRole(organizationID int, roleID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/roles/%d/users/", organizationID, roleID)
    data := map[string]interface{}{
        "id": userID,
        "disassociate": true,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// AssignTeamOrganizationRole assigns a team to a role within an organization
func (o *OrganizationsService) AssignTeamOrganizationRole(organizationID int, roleID int, teamID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/roles/%d/teams/", organizationID, roleID)
    data := map[string]interface{}{
        "id": teamID,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// RemoveTeamOrganizationRole removes a team from a role within an organization
func (o *OrganizationsService) RemoveTeamOrganizationRole(organizationID int, roleID int, teamID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/roles/%d/teams/", organizationID, roleID)
    data := map[string]interface{}{
        "id": teamID,
        "disassociate": true,
    }
    
    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}
