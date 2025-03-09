# Organization Applications

The AWX Go API provides functionality to manage applications within organizations.

## List Organization Applications

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization applications err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
apps, err := awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "name": "Terraform",
    "client_type": "confidential",
})
```

### Get All Applications
```go
// Automatically handle pagination
apps, err := awx.OrganizationsService.GetAllOrganizationApplications(9, map[string]string{
    "order_by": "name",
})
```

## Application Management

### Associate Application with Organization
```go
err := awx.OrganizationsService.AssociateApplication(9, 4)
if err != nil {
    log.Fatalf("Associate application err: %s", err)
}
```

### Disassociate Application from Organization
```go
err := awx.OrganizationsService.DisassociateApplication(9, 4)
if err != nil {
    log.Fatalf("Disassociate application err: %s", err)
}
```

## Response Types

### Application List Response
```go
type ApplicationResponse struct {
    Pagination
    Results []*Application
}

type Application struct {
    ID            int
    Type          string
    URL           string
    Related       map[string]string
    Name          string
    Description   string
    ClientID      string
    ClientType    string
    Organization  int
    Authorization_grant_type string
    Redirect_uris string
    Created       string
    Modified      string
}
```

## Error Handling

Example of proper error handling for application operations:

```go
apps, err := awx.OrganizationsService.ListOrganizationApplications(9, nil)
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

### List Applications by Type
```go
// List confidential applications
apps, err := awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "client_type": "confidential",
})

// List public applications
apps, err = awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "client_type": "public",
})
```

### Search and Sort
```go
// Search for applications
apps, err := awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "search": "terraform",
})

// Sort by name
apps, err = awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "order_by": "name",
})
```

### Filter by Grant Type
```go
// List applications by grant type
apps, err := awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "authorization_grant_type": "password",
})
```

### Application Management Example
```go
// 1. List applications
apps, err := awx.OrganizationsService.ListOrganizationApplications(9, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Find specific application
var appID int
for _, app := range apps.Results {
    if app.Name == "Terraform Integration" {
        appID = app.ID
        break
    }
}

// 3. Associate application with organization
if appID > 0 {
    err = awx.OrganizationsService.AssociateApplication(9, appID)
    if err != nil {
        log.Fatal(err)
    }
}

// 4. Verify application association
apps, err = awx.OrganizationsService.ListOrganizationApplications(9, map[string]string{
    "name": "Terraform Integration",
})
if err != nil {
    log.Fatal(err)
}
if len(apps.Results) > 0 {
    log.Printf("Application successfully associated with organization")
}
```

## Application Types and Grant Types

### Client Types
- confidential: For applications that can securely store client secrets
- public: For applications that cannot securely store client secrets

### Authorization Grant Types
- authorization-code: Standard OAuth2 authorization code flow
- password: Resource owner password credentials grant
- client-credentials: Client credentials grant
