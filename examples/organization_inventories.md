# Organization Inventories

The AWX Go API provides functionality to manage inventories within organizations.

## List Organization Inventories

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization inventories err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
inventories, err := awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "name": "Production Servers",
    "kind": "smart",
    "host_filter": "environment=prod",
})
```

### Get All Inventories
```go
// Automatically handle pagination
inventories, err := awx.OrganizationsService.GetAllOrganizationInventories(9, map[string]string{
    "order_by": "name",
})
```

## Access Control

### Associate Inventory with Organization
```go
err := awx.OrganizationsService.AssociateInventory(9, 50)
if err != nil {
    log.Fatalf("Associate inventory err: %s", err)
}
```

### Disassociate Inventory from Organization
```go
err := awx.OrganizationsService.DisassociateInventory(9, 50)
if err != nil {
    log.Fatalf("Disassociate inventory err: %s", err)
}
```

## Response Types

### Inventory List Response
```go
type InventoryResponse struct {
    Pagination
    Results []*Inventory
}

type Inventory struct {
    ID          int
    Type        string
    URL         string
    Related     map[string]string
    Name        string
    Description string
    Organization int
    Kind        string
    HostFilter  string
    Variables   string
    HasActiveFailures bool
    TotalHosts       int
    HostsWithActiveFailures int
    TotalGroups     int
    HasInventorySources bool
    // ... other fields
}
```

## Error Handling

Example of proper error handling for inventory operations:

```go
inventories, err := awx.OrganizationsService.ListOrganizationInventories(9, nil)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "404"):
        log.Fatal("Organization not found")
    case strings.Contains(err.Error(), "403"):
        log.Fatal("Permission denied")
    default:
        log.Fatal("Unknown error:", err)
    }
}
```

## Common Use Cases

### List Inventories by Type
```go
// List smart inventories
inventories, err := awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "kind": "smart",
})

// List regular inventories
inventories, err = awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "kind": "",
})
```

### List Inventories with Issues
```go
// List inventories with active failures
inventories, err := awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "has_active_failures": "true",
})

// List inventories with inventory sources
inventories, err = awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "has_inventory_sources": "true",
})
```

### Search and Sort
```go
// Search for inventories
inventories, err := awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "search": "prod",
})

// Sort by total hosts
inventories, err = awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "order_by": "-total_hosts",
})
```

### Filter by Host Counts
```go
// List inventories with more than 10 hosts
inventories, err := awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "total_hosts__gt": "10",
})

// List inventories with failing hosts
inventories, err = awx.OrganizationsService.ListOrganizationInventories(9, map[string]string{
    "hosts_with_active_failures__gt": "0",
})
