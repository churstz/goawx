# Organizations

The AWX Go API provides comprehensive functionality for managing organizations and their components.

## Core Organization Operations

### List Organizations
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizations(map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organizations err: %s", err)
}
```

### Get Organization
```go
org, err := awx.OrganizationsService.GetOrganization(9, map[string]string{})
if err != nil {
    log.Fatalf("Get organization err: %s", err)
}
```

### Create Organization
```go
result, err := awx.OrganizationsService.CreateOrganization(map[string]interface{}{
    "name": "New Organization",
    "description": "A new organization",
}, map[string]string{})
```

### Update Organization
```go
result, err := awx.OrganizationsService.UpdateOrganization(9, map[string]interface{}{
    "description": "Updated description",
}, map[string]string{})
```

### Delete Organization
```go
err := awx.OrganizationsService.DeleteOrganization(9)
if err != nil {
    log.Fatal(err)
}
```

## Organization Components

Each organization can manage multiple components. See dedicated documentation for each:

1. [Organization Users](organization_users.md)
2. [Organization Teams](organization_teams.md)
3. [Organization Admins](organization_admins.md)
4. [Organization Applications](organization_applications.md)
5. [Organization Credentials](organization_credentials.md)
6. [Organization Inventories](organization_inventories.md)
7. [Organization Job Templates](organization_job_templates.md)
8. [Organization Roles](organization_roles.md)

## Response Types

### Organization Response
```go
type OrganizationResponse struct {
    Pagination
    Results []*Organization
}

type Organization struct {
    ID          int    `json:"id"`
    Type        string `json:"type"`
    URL         string `json:"url"`
    Related     map[string]string
    Name        string `json:"name"`
    Description string `json:"description"`
    Created     string `json:"created"`
    Modified    string `json:"modified"`
}
```

## Common Use Cases

### Search Organizations
```go
// Search by name
orgs, err := awx.OrganizationsService.ListOrganizations(map[string]string{
    "name": "DevOps",
})

// Search by text
orgs, err = awx.OrganizationsService.ListOrganizations(map[string]string{
    "search": "dev",
})
```

### Sort Organizations
```go
// Sort by name ascending
orgs, err := awx.OrganizationsService.ListOrganizations(map[string]string{
    "order_by": "name",
})

// Sort by name descending
orgs, err = awx.OrganizationsService.ListOrganizations(map[string]string{
    "order_by": "-name",
})
```

### Full Organization Setup Example
```go
// 1. Create organization
org, err := awx.OrganizationsService.CreateOrganization(map[string]interface{}{
    "name": "DevOps Team",
    "description": "DevOps organization",
}, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Add admin user
err = awx.OrganizationsService.AssociateAdmin(org.ID, 24)
if err != nil {
    log.Fatal(err)
}

// 3. Create team
// See organization_teams.md for team management

// 4. Add inventory
// See organization_inventories.md for inventory management

// 5. Add job template
// See organization_job_templates.md for job template management

// 6. Setup roles
// See organization_roles.md for role management
```

## Error Handling
```go
orgs, err := awx.OrganizationsService.ListOrganizations(nil)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "401"):
        log.Fatal("Authentication failed")
    case strings.Contains(err.Error(), "403"):
        log.Fatal("Permission denied")
    case strings.Contains(err.Error(), "404"):
        log.Fatal("Resource not found")
    default:
        log.Fatal("Unknown error:", err)
    }
}
