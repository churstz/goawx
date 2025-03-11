package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// TokenService implements token management operations
type TokenService struct {
	client *Client
}

// TokenSummaryUser represents the user information in token response
type TokenSummaryUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// TokenSummaryApplication represents the application information in token response
type TokenSummaryApplication struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TokenSummaryFields represents summary fields in token response
type TokenSummaryFields struct {
	User        TokenSummaryUser        `json:"user"`
	Application TokenSummaryApplication `json:"application"`
}

// Token represents an OAuth2 token
type Token struct {
	ID            int                `json:"id"`
	Type          string             `json:"type"`
	URL           string             `json:"url"`
	Related       map[string]string  `json:"related"`
	SummaryFields TokenSummaryFields `json:"summary_fields"`
	Created       string             `json:"created"`
	Modified      string             `json:"modified"`
	Description   string             `json:"description"`
	User          int                `json:"user"`
	Token         string             `json:"token,omitempty"`
	RefreshToken  string             `json:"refresh_token,omitempty"`
	Application   int                `json:"application"`
	Expires       string             `json:"expires,omitempty"`
	Scope         string             `json:"scope"`
}

// TokenResponse represents tokens list response
type TokenResponse = PaginatedResponse[Token]

const tokenAPIEndpoint = "/api/v2/tokens/"

// ListTokens retrieves a list of all tokens
func (t *TokenService) ListTokens(params map[string]string) ([]*Token, error) {
	results, err := t.getAllPages(tokenAPIEndpoint, params)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetToken retrieves a specific token by ID
func (t *TokenService) GetToken(id int, params map[string]string) (*Token, error) {
	result := new(Token)
	endpoint := fmt.Sprintf("%s%d/", tokenAPIEndpoint, id)

	resp, err := t.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateToken updates a token
func (t *TokenService) UpdateToken(id int, data map[string]interface{}, params map[string]string) (*Token, error) {
	result := new(Token)
	endpoint := fmt.Sprintf("%s%d/", tokenAPIEndpoint, id)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// PutToken updates a token using PUT
func (t *TokenService) PutToken(id int, data map[string]interface{}, params map[string]string) (*Token, error) {
	result := new(Token)
	endpoint := fmt.Sprintf("%s%d/", tokenAPIEndpoint, id)

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Requester.PutJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteToken deletes a token
func (t *TokenService) DeleteToken(id int) error {
	endpoint := fmt.Sprintf("%s%d/", tokenAPIEndpoint, id)

	resp, err := t.client.Requester.Delete(endpoint, nil, nil)
	if err != nil {
		return err
	}

	return CheckResponse(resp)
}

func (p *TokenService) getAllPages(firstURL string, params map[string]string) ([]*Token, error) {
	results := make([]*Token, 0)
	nextURL := firstURL
	for {
		nextURLParsed, err := url.Parse(nextURL)
		if err != nil {
			return nil, err
		}

		nextURLQueryParams := make(map[string]string)
		for paramName, paramValues := range nextURLParsed.Query() {
			if len(paramValues) > 0 {
				nextURLQueryParams[paramName] = paramValues[0]
			}
		}

		for paramName, paramValue := range params {
			nextURLQueryParams[paramName] = paramValue
		}

		result := new(PaginatedResponse[Token])
		resp, err := p.client.Requester.GetJSON(nextURLParsed.Path, result, nextURLQueryParams)
		if err != nil {
			return nil, err
		}

		if err := CheckResponse(resp); err != nil {
			return nil, err
		}

		results = append(results, result.Results...)

		if result.Next == nil || result.Next.(string) == "" {
			break
		}
		nextURL = result.Next.(string)
	}
	return results, nil
}
