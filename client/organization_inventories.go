package awx

const organizationInventories = "inventories"

// OrganizationInventoryResponse represents the inventories list response
type OrganizationInventoryResponse = PaginatedResponse[Inventory]

// ListOrganizationInventories shows list of inventories in an organization.
func (o *OrganizationsService) ListOrganizationInventories(id int, params map[string]string) ([]*Inventory, error) {
	return listOrganizationResource[Inventory](o, id, organizationInventories, params)
}

// CreateOrganizationInventory creates an inventory in the specified organization.
func (o *OrganizationsService) CreateOrganizationInventory(id int, data map[string]interface{}) (*Inventory, error) {
	mandatoryFields := []string{"name"}
	return createOrganizationResource[Inventory](o, id, organizationInventories, data, mandatoryFields)
}

// AssociateInventoryWithOrganization associates an existing inventory with an organization
func (o *OrganizationsService) AssociateInventoryWithOrganization(organizationID int, inventoryID int) error {
	data := map[string]interface{}{
		"id": inventoryID,
	}

	_, err := o.associate(organizationID, organizationInventories, data, nil)
	return err
}

// DisassociateInventoryFromOrganization removes an inventory's association with an organization
func (o *OrganizationsService) DisassociateInventoryFromOrganization(organizationID int, inventoryID int) error {
	data := map[string]interface{}{
		"id": inventoryID,
	}

	_, err := o.disAssociate(organizationID, organizationInventories, data, nil)
	return err
}
