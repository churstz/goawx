# Token Management

The AWX Go API provides functionality to manage OAuth2 tokens, user tokens, and personal access tokens.

## Basic Token Operations

### Pagination and Filtering
```go
// Using pagination parameters
tokens, err := awx.Tokens.ListTokens(map[string]string{
    "page": "2",
    "page_size": "20",
})

// Using search and filtering
tokens, err := awx.Tokens.ListTokens(map[string]string{
    "application": "4",
    "scope": "write",
    "user": "24",
    "search": "terraform",
})

// Get all tokens automatically (handles pagination)
allTokens, err := awx.Tokens.ListAllTokens(nil)
```


### List All Tokens
```go
awx := awxgo.NewAWX("http://awx.your.org", "your-username", "your-password", nil)
result, err := awx.Tokens.ListTokens(map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List tokens err: %s", err)
}
```

### Get a Specific Token
```go
result, err := awx.Tokens.GetToken(5, nil)
if err != nil {
    log.Fatalf("Get token err: %s", err)
}
```

### Update a Token
```go
result, err := awx.Tokens.UpdateToken(5, map[string]interface{}{
    "description": "Updated description",
}, nil)
if err != nil {
    log.Fatalf("Update token err: %s", err)
}
```

### Delete a Token
```go
err := awx.Tokens.DeleteToken(5)
if err != nil {
    log.Fatalf("Delete token err: %s", err)
}
```

## User Token Operations

### List User Tokens
```go
// List tokens for a specific user
result, err := awx.UserService.ListUserTokens(24, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List user tokens err: %s", err)
}
```

### Get User Token
```go
result, err := awx.UserService.GetUserToken(24, 5, nil)
if err != nil {
    log.Fatalf("Get user token err: %s", err)
}
```

## Personal Access Tokens

### List Personal Access Tokens
```go
result, err := awx.UserService.ListUserPersonalTokens(24, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List personal tokens err: %s", err)
}
```

### Get Personal Access Token
```go
result, err := awx.UserService.GetUserPersonalToken(24, 5, nil)
if err != nil {
    log.Fatalf("Get personal token err: %s", err)
}
```

## Authorized Tokens

### List Authorized Tokens
```go
result, err := awx.UserService.ListUserAuthorizedTokens(24, map[string]string{
    "page": "1",
    "page_size": "10",
})
if err != nil {
    log.Fatalf("List authorized tokens err: %s", err)
}
```

### Get Authorized Token
```go
result, err := awx.UserService.GetUserAuthorizedToken(24, 5, nil)
if err != nil {
    log.Fatalf("Get authorized token err: %s", err)
}
```

## Token Types

### Token
```go
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
