# User Tokens

The AWX Go API provides functionality to manage user tokens.

## List User Tokens

### Basic Listing
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.UserService.ListUserTokens(24, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List user tokens err: %s", err)
}
```

### Using Filters
```go
// Filter by various parameters
tokens, err := awx.UserService.ListUserTokens(24, map[string]string{
    "scope": "write",
    "expires__gt": "2025-01-01T00:00:00",
})
```

### Get All Tokens
```go
// Automatically handle pagination
tokens, err := awx.UserService.GetAllUserTokens(24, map[string]string{
    "order_by": "created",
})
```

## Get Specific Token

### Get User Token
```go
token, err := awx.UserService.GetUserToken(24, 5, nil)
if err != nil {
    log.Fatalf("Get token err: %s", err)
}
```

## Response Types

### Token List Response
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
    SummaryFields map[string]interface{} `json:"summary_fields"`
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
tokens, err := awx.UserService.ListUserTokens(24, map[string]string{
    "application": "4",
})
```

### List Tokens by Scope
```go
// List write tokens
tokens, err := awx.UserService.ListUserTokens(24, map[string]string{
    "scope": "write",
})

// List read tokens
tokens, err = awx.UserService.ListUserTokens(24, map[string]string{
    "scope": "read",
})
```

### Filter by Expiration
```go
// List active tokens
tokens, err := awx.UserService.ListUserTokens(24, map[string]string{
    "expires__gt": time.Now().Format(time.RFC3339),
})

// List expired tokens
tokens, err = awx.UserService.ListUserTokens(24, map[string]string{
    "expires__lt": time.Now().Format(time.RFC3339),
})
```

### Search and Sort
```go
// Search tokens
tokens, err := awx.UserService.ListUserTokens(24, map[string]string{
    "search": "terraform",
})

// Sort by creation time
tokens, err = awx.UserService.ListUserTokens(24, map[string]string{
    "order_by": "-created",
})
```

## Error Handling
```go
tokens, err := awx.UserService.ListUserTokens(24, nil)
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

## Related Documentation

For specific token types, see:
1. [Personal Access Tokens](user_personal_tokens.md)
2. [OAuth2 Authorized Tokens](user_authorized_tokens.md)
3. [General Token Management](tokens.md)
