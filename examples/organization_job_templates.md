# Organization Job Templates

The AWX Go API provides functionality to manage job templates within organizations.

## List Organization Job Templates

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization job templates err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
templates, err := awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "name": "My Template",
    "created_by": "admin",
    "modified_by": "admin",
    "description__icontains": "deployment",
})
```

### Get All Job Templates
```go
// Automatically handle pagination
templates, err := awx.OrganizationsService.GetAllOrganizationJobTemplates(9, map[string]string{
    "order_by": "name",
})
```

## Access Control

### Associate Job Template with Organization
```go
err := awx.OrganizationsService.AssociateJobTemplate(9, 80)
if err != nil {
    log.Fatalf("Associate job template err: %s", err)
}
```

### Disassociate Job Template from Organization
```go
err := awx.OrganizationsService.DisassociateJobTemplate(9, 80)
if err != nil {
    log.Fatalf("Disassociate job template err: %s", err)
}
```

## Response Types

### Job Template List Response
```go
type JobTemplateResponse struct {
    Pagination
    Results []*JobTemplate
}

type JobTemplate struct {
    ID          int
    Type        string
    URL         string
    Related     map[string]string
    Name        string
    Description string
    // ... other fields
}
```

## Error Handling

Example of proper error handling for job template operations:

```go
templates, err := awx.OrganizationsService.ListOrganizationJobTemplates(9, nil)
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

### List Templates with Specific Criteria
```go
// List all job templates with a specific tag
templates, err := awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "job_tags": "deployment",
})

// List all job templates that run on a specific inventory
templates, err = awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "inventory": "50",
})

// List all job templates using a specific project
templates, err = awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "project": "75",
})
```

### Search Job Templates
```go
// Search for job templates with specific text in name or description
templates, err := awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "search": "deploy",
})
```

### Sort Job Templates
```go
// Sort by name ascending
templates, err := awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "order_by": "name",
})

// Sort by name descending
templates, err = awx.OrganizationsService.ListOrganizationJobTemplates(9, map[string]string{
    "order_by": "-name",
})
