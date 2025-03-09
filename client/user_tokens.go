package awx

import (
    "fmt"
)

// TokenResponse represents the tokens list response
type TokenResponse struct {
    Pagination
    Results []*Token `json:"results"`
}

// ListUserTokens shows list of regular OAuth2 tokens for a user
// The returned tokens can be managed using the TokenService
func (u *UserService) ListUserTokens(id int, params map[string]string) ([]*Token, error) {
    result := new(TokenResponse)
    endpoint := fmt.Sprintf("/api/v2/users/%d/tokens/", id)
    
    resp, err := u.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// GetUserToken retrieves a specific token via the TokenService
func (u *UserService) GetUserToken(userID int, tokenID int, params map[string]string) (*Token, error) {
    // Use TokenService for actual token operations since the URLs
    // from ListUserTokens point to /api/v2/tokens/{id}/
    return u.client.Tokens.GetToken(tokenID, params)
}
