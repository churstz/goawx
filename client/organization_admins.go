package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationAdminResponse represents the admins list response
type OrganizationAdminResponse struct {
    Pagination
    Results []*User `json:"results"` // Using the base User type since these are user objects
}

// ListOrganizationAdmins shows list of admin users for an organization.
func (p *OrganizationsService) ListOrganizationAdmins(id int, params map[string]string) ([]*User, error) {
    result := new(OrganizationAdminResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/admins/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// AssociateOrganizationAdmin makes a user an admin of the organization
func (p *OrganizationsService) AssociateOrganizationAdmin(organizationID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/admins/", organizationID)
    data := map[string]interface{}{
        "id": userID,
        "associate": true,
    }

    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// DisassociateOrganizationAdmin removes a user's admin role from the organization
func (p *OrganizationsService) DisassociateOrganizationAdmin(organizationID int, userID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/admins/", organizationID)
    data := map[string]interface{}{
        "id": userID,
        "disassociate": true,
    }

    payload, err := json.Marshal(data)
    if err != nil {
        return err
    }

    resp, err := p.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// IsOrganizationAdmin checks if a specific user is an admin of the organization
func (p *OrganizationsService) IsOrganizationAdmin(organizationID int, userID int) (bool, error) {
    admins, err := p.ListOrganizationAdmins(organizationID, nil)
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
