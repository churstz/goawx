package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

// OrganizationInventoryResponse represents the inventories list response
type OrganizationInventoryResponse struct {
    Pagination
    Results []*Inventory `json:"results"` // Using the base Inventory type since the response matches
}

// ListOrganizationInventories shows list of inventories in an organization.
func (p *OrganizationsService) ListOrganizationInventories(id int, params map[string]string) ([]*Inventory, error) {
    result := new(OrganizationInventoryResponse)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/inventories/", id)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// CreateOrganizationInventory creates an inventory in the specified organization.
func (p *OrganizationsService) CreateOrganizationInventory(id int, data map[string]interface{}, params map[string]string) (*Inventory, error) {
    mandatoryFields := []string{"name"}
    validate, status := ValidateParams(data, mandatoryFields)
    if !status {
        err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    // Ensure the organization field is set to the correct ID
    data["organization"] = id

    result := new(Inventory)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/inventories/", id)
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

// GetOrganizationInventory retrieves a specific inventory in the organization
func (p *OrganizationsService) GetOrganizationInventory(organizationID int, inventoryID int, params map[string]string) (*Inventory, error) {
    result := new(Inventory)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/inventories/%d/", organizationID, inventoryID)
    
    resp, err := p.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateOrganizationInventory updates an inventory in the organization
func (p *OrganizationsService) UpdateOrganizationInventory(organizationID int, inventoryID int, data map[string]interface{}, params map[string]string) (*Inventory, error) {
    result := new(Inventory)
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/inventories/%d/", organizationID, inventoryID)
    
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

// DeleteOrganizationInventory removes an inventory from the organization
func (p *OrganizationsService) DeleteOrganizationInventory(organizationID int, inventoryID int) error {
    endpoint := fmt.Sprintf("/api/v2/organizations/%d/inventories/%d/", organizationID, inventoryID)
    
    resp, err := p.client.Requester.Delete(endpoint, nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}
