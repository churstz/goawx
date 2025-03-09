# Role Management

The AWX Go API provides comprehensive role management functionality for various resources.

## Basic Role Operations

### Pagination and Filtering
```go
// Using pagination parameters
roles, err := awx.Roles.ListRoles(map[string]string{
    "page": "2",
    "page_size": "20",
})

// Using search and filtering
roles, err := awx.Roles.ListRoles(map[string]string{
    "name": "Admin",
    "resource_type": "project",
    "search": "template",
})
```


### List All Roles
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)

// List roles with pagination
result, err := awx.Roles.ListRoles(map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List roles err: %s", err)
}

// Get all roles automatically (handles pagination)
allRoles, err := awx.Roles.ListAllRoles(nil)
if err != nil {
    log.Fatalf("List all roles err: %s", err)
}
```

### Get a Specific Role
```go
role, err := awx.Roles.GetRole(507, nil)
if err != nil {
    log.Fatalf("Get role err: %s", err)
}
```

## Resource Role Operations

### Project Roles
```go
// List roles for a project with pagination
roles, err := awx.Roles.GetProjectRoles(79, map[string]string{
    "page": "1",
})

// List all project roles (auto-pagination)
allRoles, err := awx.Roles.GetAllProjectRoles(79, nil)

// Get project object roles
objectRoles, err := awx.Roles.GetProjectObjectRoles(79, nil)
```

### Job Template Roles
```go
// List roles for a job template with pagination
roles, err := awx.Roles.GetJobTemplateRoles(80, map[string]string{
    "page": "1",
})

// List all job template roles (auto-pagination)
allRoles, err := awx.Roles.GetAllJobTemplateRoles(80, nil)

// Get job template object roles
objectRoles, err := awx.Roles.GetJobTemplateObjectRoles(80, nil)
```

### Inventory Roles
```go
// List roles for an inventory with pagination
roles, err := awx.Roles.GetInventoryRoles(50, map[string]string{
    "page": "1",
})

// List all inventory roles (auto-pagination)
allRoles, err := awx.Roles.GetAllInventoryRoles(50, nil)

// Get inventory object roles
objectRoles, err := awx.Roles.GetInventoryObjectRoles(50, nil)
```

## Role Assignment Operations

### Manage User Roles
```go
// Assign user to role
err := awx.Roles.AssignUserRole(507, 24)
if err != nil {
    log.Fatalf("Assign user role err: %s", err)
}

// Remove user from role
err = awx.Roles.RemoveUserRole(507, 24)
if err != nil {
    log.Fatalf("Remove user role err: %s", err)
}

// List users in a role
users, err := awx.Roles.ListRoleUsers(507, nil)
if err != nil {
    log.Fatalf("List role users err: %s", err)
}
```

### Manage Team Roles
```go
// Assign team to role
err := awx.Roles.AssignTeamRole(507, 15)
if err != nil {
    log.Fatalf("Assign team role err: %s", err)
}

// Remove team from role
err = awx.Roles.RemoveTeamRole(507, 15)
if err != nil {
    log.Fatalf("Remove team role err: %s", err)
}

// List teams in a role
teams, err := awx.Roles.ListRoleTeams(507, nil)
if err != nil {
    log.Fatalf("List role teams err: %s", err)
}
```

## Organization Role Operations

### List Organization Roles
```go
// List roles for an organization
roles, err := awx.OrganizationsService.ListOrganizationRoles(9, map[string]string{
    "page": "1",
})

// List object roles for an organization
objectRoles, err := awx.OrganizationsService.ListOrganizationObjectRoles(9, nil)
```

### Manage Organization Role Assignments
```go
// Assign user to organization role
err := awx.OrganizationsService.AssignUserOrganizationRole(9, 507, 24)
if err != nil {
    log.Fatalf("Assign user org role err: %s", err)
}

// Remove user from organization role
err = awx.OrganizationsService.RemoveUserOrganizationRole(9, 507, 24)
if err != nil {
    log.Fatalf("Remove user org role err: %s", err)
}

// Assign team to organization role
err := awx.OrganizationsService.AssignTeamOrganizationRole(9, 507, 15)
if err != nil {
    log.Fatalf("Assign team org role err: %s", err)
}

// Remove team from organization role
err = awx.OrganizationsService.RemoveTeamOrganizationRole(9, 507, 15)
if err != nil {
    log.Fatalf("Remove team org role err: %s", err)
}
```

## Role Types

### Role
```go
type Role struct {
    ID            int
    Type          string
    URL           string
    Related       map[string]string
    SummaryFields RoleSummaryFields
    Name          string
    Description   string
}

type RoleSummaryFields struct {
    ResourceName           string
    ResourceType          string
    ResourceTypeDisplayName string
    ResourceID           int
}
