# Organization Administrators

The AWX Go API provides functionality to manage administrators within organizations.

## List Organization Admins

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization admins err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
admins, err := awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "username": "admin",
    "is_superuser": "true",
})
```

### Get All Admins
```go
// Automatically handle pagination
admins, err := awx.OrganizationsService.GetAllOrganizationAdmins(9, map[string]string{
    "order_by": "username",
})
```

## Admin Management

### Associate Admin with Organization
```go
err := awx.OrganizationsService.AssociateAdmin(9, 24)
if err != nil {
    log.Fatalf("Associate admin err: %s", err)
}
```

### Disassociate Admin from Organization
```go
err := awx.OrganizationsService.DisassociateAdmin(9, 24)
if err != nil {
    log.Fatalf("Disassociate admin err: %s", err)
}
```

## Response Types

### Admin List Response
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
    Created         string
    Modified        string
}
```

## Error Handling

Example of proper error handling for admin operations:

```go
admins, err := awx.OrganizationsService.ListOrganizationAdmins(9, nil)
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

### List Admins by Status
```go
// List superuser admins
admins, err := awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "is_superuser": "true",
})

// List system auditor admins
admins, err = awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "is_system_auditor": "true",
})
```

### Search and Sort
```go
// Search for admins
admins, err := awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "search": "john",
})

// Sort by username
admins, err = awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "order_by": "username",
})
```

### Filter by Last Login
```go
// List recently active admins
admins, err := awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "last_login__gt": "2025-01-01T00:00:00",
})
```

### Admin Management Example
```go
// 1. List admins
admins, err := awx.OrganizationsService.ListOrganizationAdmins(9, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Find specific admin
var adminID int
for _, admin := range admins.Results {
    if admin.Username == "jdoe" {
        adminID = admin.ID
        break
    }
}

// 3. Associate admin with organization
if adminID > 0 {
    err = awx.OrganizationsService.AssociateAdmin(9, adminID)
    if err != nil {
        log.Fatal(err)
    }
}

// 4. Verify admin access
admins, err = awx.OrganizationsService.ListOrganizationAdmins(9, map[string]string{
    "username": "jdoe",
})
if err != nil {
    log.Fatal(err)
}
if len(admins.Results) > 0 {
    log.Printf("Admin successfully associated with organization")
}
