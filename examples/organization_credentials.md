# Organization Credentials

The AWX Go API provides functionality to manage credentials within organizations.

## List Organization Credentials

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List organization credentials err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
creds, err := awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "name": "AWS Access",
    "credential_type": "aws",
})
```

### Get All Credentials
```go
// Automatically handle pagination
creds, err := awx.OrganizationsService.GetAllOrganizationCredentials(9, map[string]string{
    "order_by": "name",
})
```

## Credential Management

### Associate Credential with Organization
```go
err := awx.OrganizationsService.AssociateCredential(9, 26)
if err != nil {
    log.Fatalf("Associate credential err: %s", err)
}
```

### Disassociate Credential from Organization
```go
err := awx.OrganizationsService.DisassociateCredential(9, 26)
if err != nil {
    log.Fatalf("Disassociate credential err: %s", err)
}
```

## Response Types

### Credential List Response
```go
type CredentialResponse struct {
    Pagination
    Results []*Credential
}

type Credential struct {
    ID            int
    Type          string
    URL           string
    Related       map[string]string
    SummaryFields map[string]interface{} `json:"summary_fields"`
    Name          string
    Description   string
    Organization  int
    CredentialType int    `json:"credential_type"`
    Created       string
    Modified      string
    Inputs        map[string]interface{}
}
```

## Error Handling

Example of proper error handling for credential operations:

```go
creds, err := awx.OrganizationsService.ListOrganizationCredentials(9, nil)
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

### List Credentials by Type
```go
// List AWS credentials
creds, err := awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "credential_type__name": "Amazon Web Services",
})

// List SSH credentials
creds, err = awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "credential_type__name": "Machine",
})
```

### Search and Sort
```go
// Search for credentials
creds, err := awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "search": "aws",
})

// Sort by name
creds, err = awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "order_by": "name",
})
```

### Filter by Creation Time
```go
// List recently created credentials
creds, err := awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "created__gt": "2025-01-01T00:00:00",
})
```

### Credential Management Example
```go
// 1. List credentials
creds, err := awx.OrganizationsService.ListOrganizationCredentials(9, nil)
if err != nil {
    log.Fatal(err)
}

// 2. Find specific credential
var credID int
for _, cred := range creds.Results {
    if cred.Name == "AWS Production" {
        credID = cred.ID
        break
    }
}

// 3. Associate credential with organization
if credID > 0 {
    err = awx.OrganizationsService.AssociateCredential(9, credID)
    if err != nil {
        log.Fatal(err)
    }
}

// 4. Verify credential association
creds, err = awx.OrganizationsService.ListOrganizationCredentials(9, map[string]string{
    "name": "AWS Production",
})
if err != nil {
    log.Fatal(err)
}
if len(creds.Results) > 0 {
    log.Printf("Credential successfully associated with organization")
}
```

## Common Credential Types
- Machine: SSH/Machine credentials
- Source Control: Git and other SCM credentials
- Vault: HashiCorp Vault credentials
- AWS: Amazon Web Services credentials
- Azure: Microsoft Azure credentials
- GCP: Google Cloud Platform credentials
- OpenStack: OpenStack credentials
- VMware: VMware vCenter credentials
- Red Hat Satellite 6: Satellite credentials
- Red Hat Insights: Insights credentials
- Ansible Tower: Tower/AWX credentials
