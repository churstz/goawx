# Organization Users

The AWX Go API provides functionality to manage users within organizations.

## List Organization Users

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization users err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "username": "jdoe",
    "is_superuser": "false",
})
```

### Get All Users
```go
// Automatically handle pagination
users, err := awx.OrganizationsService.GetAllOrganizationUsers(9, map[string]string{
    "order_by": "username",
})
```

## User Management

### Associate User with Organization
```go
err := awx.OrganizationsService.AssociateUser(9, 24)
if err != nil {
    log.Fatalf("Associate user err: %s", err)
}
```

### Disassociate User from Organization
```go
err := awx.OrganizationsService.DisassociateUser(9, 24)
if err != nil {
    log.Fatalf("Disassociate user err: %s", err)
}
```

## Response Types

### User List Response
```go
type UserResponse struct {
    Pagination
    Results []*User
}

type User struct {
    ID              int
    Type            string
    URL             string
    Related         map[string]string
    Username        string
    FirstName       string
    LastName        string
    Email           string
    IsSuperuser     bool
    IsSystemAuditor bool
    Password        string `json:"password,omitempty"`
    Created         string
    Modified        string
}
```

## Error Handling

Example of proper error handling for user operations:

```go
users, err := awx.OrganizationsService.ListOrganizationUsers(9, nil)
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

### List Users by Type
```go
// List superusers
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "is_superuser": "true",
})

// List system auditors
users, err = awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "is_system_auditor": "true",
})
```

### Search and Sort
```go
// Search for users
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "search": "john",
})

// Sort by username
users, err = awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "order_by": "username",
})
```

### Filter by Activity
```go
// List recently active users
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "last_login__gt": "2025-01-01T00:00:00",
})

// List recently modified users
users, err = awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "modified__gt": "2025-01-01T00:00:00",
})
```

### User Management Example
```go
// 1. List users
users, err := awx.OrganizationsService.ListOrganizationUsers(9, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Find specific user
var userID int
for _, user := range users.Results {
    if user.Username == "jdoe" {
        userID = user.ID
        break
    }
}

// 3. Associate user with organization
if userID > 0 {
    err = awx.OrganizationsService.AssociateUser(9, userID)
    if err != nil {
        log.Fatal(err)
    }
}

// 4. Verify user association
users, err = awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "username": "jdoe",
})
if err != nil {
    log.Fatal(err)
}
if len(users.Results) > 0 {
    log.Printf("User successfully associated with organization")
}
```

## Special Filters

### Filter by Teams
```go
// List users in specific team
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "team": "15",
})
```

### Filter by Roles
```go
// List users with specific role
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "role": "admin_role",
})
```

### Filter by Permissions
```go
// List users with specific permission
users, err := awx.OrganizationsService.ListOrganizationUsers(9, map[string]string{
    "has_permissions": "inventory.admin",
})
