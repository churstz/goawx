package awx

import (
    "fmt"
)

// AuthorizedToken represents an OAuth2 token authorized by a user
type AuthorizedToken struct {
    ID            int               `json:"id"`
    Type          string           `json:"type"`
    URL           string           `json:"url"`
    Related       map[string]string `json:"related"`
    Created       string           `json:"created"`
    Modified      string           `json:"modified"`
    Description   string           `json:"description"`
    User          int             `json:"user"`
    Token         string          `json:"token,omitempty"`
    Application   int             `json:"application"`
    Scope         string          `json:"scope"`
    Expires       string          `json:"expires,omitempty"`
    RefreshToken  string          `json:"refresh_token,omitempty"`
}

// AuthorizedTokenResponse represents the tokens list response
type AuthorizedTokenListResponse struct {
    Pagination
    Results []*AuthorizedToken `json:"results"`
}

// ListUserAuthorizedTokens shows list of tokens authorized by the user
func (u *UserService) ListUserAuthorizedTokens(id int, params map[string]string) ([]*AuthorizedToken, error) {
    result := new(AuthorizedTokenListResponse)
    endpoint := fmt.Sprintf("/api/v2/users/%d/authorized_tokens/", id)
    
    resp, err := u.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// GetUserAuthorizedToken retrieves a specific authorized token
func (u *UserService) GetUserAuthorizedToken(userID int, tokenID int, params map[string]string) (*AuthorizedToken, error) {
    result := new(AuthorizedToken)
    endpoint := fmt.Sprintf("/api/v2/users/%d/authorized_tokens/%d/", userID, tokenID)
    
    resp, err := u.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}
