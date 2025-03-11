package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const organizationInventoriesAPIEndpoint = "/api/v2/organizations/%d/inventories/"

// OrganizationInventoryResponse represents the inventories list response
type OrganizationInventoryResponse = PaginatedResponse[Inventory]

// ListOrganizationInventories shows list of inventories in an organization.
func (o *OrganizationsService) ListOrganizationInventories(id int, params map[string]string) ([]*Inventory, error) {
	result := new(OrganizationInventoryResponse)
	endpoint := fmt.Sprintf(organizationInventoriesAPIEndpoint, id)

	resp, err := o.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// CreateOrganizationInventory creates an inventory in the specified organization.
func (o *OrganizationsService) CreateOrganizationInventory(id int, data map[string]interface{}, params map[string]string) (*Inventory, error) {
	mandatoryFields := []string{"name"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("mandatory input arguments are absent: %s", validate)
		return nil, err
	}

	// Ensure the organization field is set to the correct ID
	data["organization"] = id

	result := new(Inventory)
	endpoint := fmt.Sprintf(organizationInventoriesAPIEndpoint, id)
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

// AssociateInventoryWithOrganization associates an existing inventory with an organization
func (o *OrganizationsService) AssociateInventoryWithOrganization(organizationID int, inventoryID int) error {
	data := map[string]interface{}{
		"id": inventoryID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationInventoriesAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}

// DisassociateInventoryFromOrganization removes an inventory's association with an organization
func (o *OrganizationsService) DisassociateInventoryFromOrganization(organizationID int, inventoryID int) error {
	data := map[string]interface{}{
		"id":           inventoryID,
		"disassociate": true,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf(organizationInventoriesAPIEndpoint, organizationID)
	resp, err := o.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}
