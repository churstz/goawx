# Organization Teams

The AWX Go API provides functionality to manage teams within organizations.

## List Organization Teams

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization teams err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
teams, err := awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "name": "DevOps",
    "description__icontains": "automation",
})
```

### Get All Teams
```go
// Automatically handle pagination
teams, err := awx.OrganizationsService.GetAllOrganizationTeams(9, map[string]string{
    "order_by": "name",
})
```

## Access Control

### Associate Team with Organization
```go
err := awx.OrganizationsService.AssociateTeam(9, 15)
if err != nil {
    log.Fatalf("Associate team err: %s", err)
}
```

### Disassociate Team from Organization
```go
err := awx.OrganizationsService.DisassociateTeam(9, 15)
if err != nil {
    log.Fatalf("Disassociate team err: %s", err)
}
```

## Response Types

### Team List Response
```go
type TeamResponse struct {
    Pagination
    Results []*Team
}

type Team struct {
    ID          int
    Type        string
    URL         string
    Related     map[string]string
    Name        string
    Description string
    Organization int
    Summary_fields map[string]interface{}
    Created     string
    Modified    string
}
```

## Error Handling

Example of proper error handling for team operations:

```go
teams, err := awx.OrganizationsService.ListOrganizationTeams(9, nil)
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

### List Teams by Roles
```go
// List teams with admin access
teams, err := awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "role": "admin_role",
})

// List teams with use access
teams, err = awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "role": "use_role",
})
```

### Search and Sort
```go
// Search for teams
teams, err := awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "search": "ops",
})

// Sort by name
teams, err = awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "order_by": "name",
})
```

### Filter by Creation/Modification Time
```go
// List recently created teams
teams, err := awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "created__gt": "2025-01-01T00:00:00",
})

// List recently modified teams
teams, err = awx.OrganizationsService.ListOrganizationTeams(9, map[string]string{
    "modified__gt": "2025-03-01T00:00:00",
})
```

### Team Management Example
```go
// 1. List teams
teams, err := awx.OrganizationsService.ListOrganizationTeams(9, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Find specific team
var teamID int
for _, team := range teams.Results {
    if team.Name == "DevOps" {
        teamID = team.ID
        break
    }
}

// 3. Associate team with organization
if teamID > 0 {
    err = awx.OrganizationsService.AssociateTeam(9, teamID)
    if err != nil {
        log.Fatal(err)
    }
}
