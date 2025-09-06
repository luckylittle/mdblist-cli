package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiBaseURL = "https://api.mdblist.com"
)

// Client holds the http client and api key for making requests to MDBList API.
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// New creates a new MDBList API client.
func New(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("MDBList API key is required")
	}
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

// GetMyLimits fetches the current user's API limits.
func (c *Client) GetMyLimits() (interface{}, error) {
	return c.doRequest("/user")
}

// GetMyLists fetches the current user's lists.
func (c *Client) GetMyLists() (interface{}, error) {
	return c.doRequest("/lists/user")
}

func (c *Client) doRequest(endpoint string) (interface{}, error) {
	url := fmt.Sprintf("%s%s?apikey=%s", apiBaseURL, endpoint, c.apiKey)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	// The API key is now in the query parameters.
	// req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		// If it's not JSON, it might be an error string. Return the raw body.
		return string(body), fmt.Errorf("failed to unmarshal JSON response (status %d): %s", resp.StatusCode, string(body))
	}

	if resp.StatusCode >= 400 {
		return result, fmt.Errorf("API error (status %d)", resp.StatusCode)
	}

	return result, nil
}
