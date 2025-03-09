package awx

import (
    "fmt"
)

// PersonalToken represents a personal access token
type PersonalToken struct {
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
}

// PersonalTokenListResponse represents the personal tokens list response
type PersonalTokenListResponse struct {
    Pagination
    Results []*PersonalToken `json:"results"`
}

// ListUserPersonalTokens shows list of user's personal tokens
func (u *UserService) ListUserPersonalTokens(id int, params map[string]string) ([]*PersonalToken, error) {
    result := new(PersonalTokenListResponse)
    endpoint := fmt.Sprintf("/api/v2/users/%d/personal_tokens/", id)
    
    resp, err := u.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}

// GetUserPersonalToken retrieves a specific personal token
func (u *UserService) GetUserPersonalToken(userID int, tokenID int, params map[string]string) (*PersonalToken, error) {
    result := new(PersonalToken)
    endpoint := fmt.Sprintf("/api/v2/users/%d/personal_tokens/%d/", userID, tokenID)
    
    resp, err := u.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}
