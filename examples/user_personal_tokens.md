# User Personal Access Tokens

The AWX Go API provides functionality to manage personal access tokens (PAT) for users.

## List Personal Access Tokens

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List personal tokens err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
tokens, err := awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "description__icontains": "terraform",
    "scope": "write",
})
```

### Get All Personal Tokens
```go
// Automatically handle pagination
tokens, err := awx.UserService.GetAllUserPersonalTokens(24, map[string]string{
    "order_by": "created",
})
```

## Get Specific Personal Token

### Get Token Details
```go
token, err := awx.UserService.GetUserPersonalToken(24, 5, nil)
if err != nil {
    log.Fatalf("Get personal token err: %s", err)
}
```

## Response Types

### Personal Token Response
```go
type TokenResponse struct {
    Pagination
    Results []*Token
}

type Token struct {
    ID            int
    Type          string
    URL           string
    Related       map[string]string
    Created       string
    Modified      string
    Description   string
    User          int
    Token         string          // Only returned when creating token
    Scope         string
    Expires       string
}
```

## Common Use Cases

### List Tokens by Scope
```go
// List write tokens
tokens, err := awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "scope": "write",
})

// List read tokens
tokens, err = awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "scope": "read",
})
```

### Filter by Expiration
```go
// List active tokens
tokens, err := awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "expires__gt": time.Now().Format(time.RFC3339),
})

// List expired tokens
tokens, err = awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "expires__lt": time.Now().Format(time.RFC3339),
})
```

### Search and Sort
```go
// Search tokens
tokens, err := awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "search": "automation",
})

// Sort by creation time
tokens, err = awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "order_by": "-created",
})
```

## Error Handling
```go
tokens, err := awx.UserService.ListUserPersonalTokens(24, nil)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "401"):
        log.Fatal("Authentication failed")
    case strings.Contains(err.Error(), "403"):
        log.Fatal("Permission denied")
    case strings.Contains(err.Error(), "404"):
        log.Fatal("User not found")
    default:
        log.Fatal("Unknown error:", err)
    }
}
```

## Personal Access Token Best Practices

1. Scoping:
   - Use minimal required scope (read vs. write)
   - Consider using multiple tokens for different purposes

2. Expiration:
   - Set appropriate expiration dates
   - Regularly rotate tokens
   - Remove expired or unused tokens

3. Description:
   - Use descriptive names
   - Include purpose or application name
   - Note environment (dev/prod)

## Related Documentation

For other token types, see:
1. [User Tokens](user_tokens.md)
2. [OAuth2 Authorized Tokens](user_authorized_tokens.md)
3. [General Token Management](tokens.md)
