# Organization Role Management

The AWX Go API provides functionality to manage roles within organizations.

## List Organization Roles

### List All Organization Roles
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationRoles(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization roles err: %s", err)
}
```

### List Organization Object Roles
```go
result, err := awx.OrganizationsService.ListOrganizationObjectRoles(9, nil)
if err != nil {
    log.Fatalf("List organization object roles err: %s", err)
}
```

## Role Assignment Management

### User Role Management
```go
// Assign a user to an organization role
err := awx.OrganizationsService.AssignUserOrganizationRole(9, 507, 24)
if err != nil {
    log.Fatalf("Assign user role err: %s", err)
}

// Remove a user from an organization role
err = awx.OrganizationsService.RemoveUserOrganizationRole(9, 507, 24)
if err != nil {
    log.Fatalf("Remove user role err: %s", err)
}

// List users for a specific organization role
users, err := awx.OrganizationsService.ListOrganizationRoleUsers(9, 507, map[string]string{
    "page": "1",
    "page_size": "20",
})
```

### Team Role Management
```go
// Assign a team to an organization role
err := awx.OrganizationsService.AssignTeamOrganizationRole(9, 507, 15)
if err != nil {
    log.Fatalf("Assign team role err: %s", err)
}

// Remove a team from an organization role
err = awx.OrganizationsService.RemoveTeamOrganizationRole(9, 507, 15)
if err != nil {
    log.Fatalf("Remove team role err: %s", err)
}

// List teams for a specific organization role
teams, err := awx.OrganizationsService.ListOrganizationRoleTeams(9, 507, map[string]string{
    "page": "1",
    "page_size": "20",
})
```

## Filtering and Searching

### Using Search Parameters
```go
roles, err := awx.OrganizationsService.ListOrganizationRoles(9, map[string]string{
    "search": "admin",
    "role_type": "admin_role",
})
```

### Using Pagination
```go
roles, err := awx.OrganizationsService.ListOrganizationRoles(9, map[string]string{
    "page": "2",
    "page_size": "25",
    "order_by": "name",
})
```

## Common Role Types

Organization roles typically include:
- Admin Role: Full administrative access
- Execute Role: Can execute jobs within the organization
- Project Admin Role: Can manage projects
- Inventory Admin Role: Can manage inventories
- Credential Admin Role: Can manage credentials
- Read Role: Read-only access to the organization

## Example: Full Organization Role Setup
```go
// 1. List available roles
roles, err := awx.OrganizationsService.ListOrganizationRoles(9, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Find the admin role
var adminRoleID int
for _, role := range roles.Results {
    if role.Name == "Admin" {
        adminRoleID = role.ID
        break
    }
}

// 3. Assign a user as organization admin
err = awx.OrganizationsService.AssignUserOrganizationRole(9, adminRoleID, 24)
if err != nil {
    log.Fatal(err)
}

// 4. Assign a team as organization admin
err = awx.OrganizationsService.AssignTeamOrganizationRole(9, adminRoleID, 15)
if err != nil {
    log.Fatal(err)
}
```

## Error Handling

Example of proper error handling when working with organization roles:

```go
roles, err := awx.OrganizationsService.ListOrganizationRoles(9, nil)
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
