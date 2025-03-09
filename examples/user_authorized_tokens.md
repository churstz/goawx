# User Authorized Tokens

The AWX Go API provides functionality to manage OAuth2 authorized tokens for users.

## List Authorized Tokens

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List authorized tokens err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
tokens, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "application": "4",
    "scope": "write",
})
```

### Get All Authorized Tokens
```go
// Automatically handle pagination
tokens, err := awx.UserService.GetAllUserAuthorizedTokens(24, map[string]string{
    "order_by": "created",
})
```

## Get Specific Authorized Token

### Get Token Details
```go
token, err := awx.UserService.GetUserAuthorizedToken(24, 5, nil)
if err != nil {
    log.Fatalf("Get authorized token err: %s", err)
}
```

## Response Types

### Authorized Token Response
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
    Application   int
    Scope         string
    Expires       string
    RefreshToken  string          // Only returned when creating token
}
```

## Common Use Cases

### List Tokens by Application
```go
// List tokens for specific application
tokens, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "application": "4",
})

// List tokens for multiple applications
tokens, err = awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "application__in": "4,5,6",
})
```

### List Tokens by Scope
```go
// List write tokens
tokens, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "scope": "write",
})

// List read tokens
tokens, err = awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "scope": "read",
})
```

### Filter by Expiration
```go
// List active tokens
tokens, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "expires__gt": time.Now().Format(time.RFC3339),
})

// List expired tokens
tokens, err = awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "expires__lt": time.Now().Format(time.RFC3339),
})
```

### Search and Sort
```go
// Search tokens
tokens, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "search": "terraform",
})

// Sort by creation time
tokens, err = awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "order_by": "-created",
})
```

## Error Handling
```go
tokens, err := awx.UserService.ListUserAuthorizedTokens(24, nil)
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

## OAuth2 Authorization Best Practices

1. Application Management:
   - Regularly review authorized applications
   - Remove access for unused applications
   - Monitor application-specific token usage

2. Token Lifecycle:
   - Implement token refresh logic
   - Handle token expiration gracefully
   - Revoke tokens when no longer needed

3. Security:
   - Store refresh tokens securely
   - Never log token values
   - Use HTTPS for all token operations

## Related Documentation

For other token types, see:
1. [User Tokens](user_tokens.md)
2. [Personal Access Tokens](user_personal_tokens.md)
3. [General Token Management](tokens.md)
