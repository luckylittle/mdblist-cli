package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
func (c *Client) GetMyLimits() (*MyLimits, error) {
	var limits MyLimits
	err := c.doRequest(http.MethodGet, "/user", nil, nil, &limits)
	return &limits, err
}

// GetMyLists fetches the current user's lists.
func (c *Client) GetMyLists() ([]List, error) {
	var lists []List
	err := c.doRequest(http.MethodGet, "/lists/user", nil, nil, &lists)
	return lists, err
}

// GetUserListsByID fetches a user's lists by their user ID.
func (c *Client) GetUserListsByID(userID int) ([]List, error) {
	endpoint := fmt.Sprintf("/lists/user/%d", userID)
	var lists []List
	err := c.doRequest(http.MethodGet, endpoint, nil, nil, &lists)
	return lists, err
}

// GetUserListsByName fetches a user's lists by their username.
func (c *Client) GetUserListsByName(username string) ([]List, error) {
	endpoint := fmt.Sprintf("/lists/user/%s", username)
	var lists []List
	err := c.doRequest(http.MethodGet, endpoint, nil, nil, &lists)
	return lists, err
}

// GetListByID fetches a list by its ID.
func (c *Client) GetListByID(listID int) ([]List, error) {
	endpoint := fmt.Sprintf("/lists/%d", listID)
	var list []List
	err := c.doRequest(http.MethodGet, endpoint, nil, nil, &list)
	return list, err
}

// UpdateListNameByID updates a list's name by its ID.
func (c *Client) UpdateListNameByID(listID int, newName string) (*ListUpdateResponse, error) {
	endpoint := fmt.Sprintf("/lists/%d", listID)
	payload := map[string]string{"name": newName}
	var response ListUpdateResponse
	err := c.doRequest(http.MethodPut, endpoint, nil, payload, &response)
	return &response, err
}

// GetListByName fetches a list by username and list name (slug).
func (c *Client) GetListByName(username, listname string) ([]List, error) {
	endpoint := fmt.Sprintf("/lists/%s/%s", username, listname)
	var lists []List
	err := c.doRequest(http.MethodGet, endpoint, nil, nil, &lists)
	return lists, err
}

// UpdateListNameByName updates a list's name by username and list name.
func (c *Client) UpdateListNameByName(username, listname, newName string) (*ListUpdateResponse, error) {
	endpoint := fmt.Sprintf("/lists/%s/%s", username, listname)
	payload := map[string]string{"name": newName}
	var response ListUpdateResponse
	err := c.doRequest(http.MethodPut, endpoint, nil, payload, &response)
	return &response, err
}

// GetListItems fetches items from a list by its ID.
func (c *Client) GetListItems(listID int, params url.Values) (*ListItems, error) {
	endpoint := fmt.Sprintf("/lists/%d/items", listID)
	var items ListItems
	err := c.doRequest(http.MethodGet, endpoint, params, nil, &items)
	return &items, err
}

// GetListItemsByName fetches items from a list by username and list name.
func (c *Client) GetListItemsByName(username, listname string, params url.Values) (*ListItems, error) {
	endpoint := fmt.Sprintf("/lists/%s/%s/items", username, listname)
	var items ListItems
	err := c.doRequest(http.MethodGet, endpoint, params, nil, &items)
	return &items, err
}

// GetListChanges fetches changes for a list by its ID.
func (c *Client) GetListChanges(listID int) (*ListChanges, error) {
	endpoint := fmt.Sprintf("/lists/%d/changes", listID)
	var changes ListChanges
	err := c.doRequest(http.MethodGet, endpoint, nil, nil, &changes)
	return &changes, err
}

// GetMediaInfo fetches information about a media item.
func (c *Client) GetMediaInfo(provider, mediaType, mediaID string, params url.Values) (*MediaInfo, error) {
	endpoint := fmt.Sprintf("/%s/%s/%s", provider, mediaType, mediaID)
	var info MediaInfo
	err := c.doRequest(http.MethodGet, endpoint, params, nil, &info)
	return &info, err
}

// GetMediaInfoBatch fetches information for multiple media items.
func (c *Client) GetMediaInfoBatch(provider, mediaType string, body MediaInfoBatchRequest) ([]MediaInfo, error) {
	endpoint := fmt.Sprintf("/%s/%s", provider, mediaType)
	var info []MediaInfo
	err := c.doRequest(http.MethodPost, endpoint, nil, body, &info)
	return info, err
}

// SearchMedia searches for media.
func (c *Client) SearchMedia(mediaType string, params url.Values) (*SearchResult, error) {
	endpoint := fmt.Sprintf("/search/%s", mediaType)
	var result SearchResult
	err := c.doRequest(http.MethodGet, endpoint, params, nil, &result)
	return &result, err
}

// GetTopLists fetches the top lists.
func (c *Client) GetTopLists() ([]List, error) {
	var lists []List
	err := c.doRequest(http.MethodGet, "/lists/top", nil, nil, &lists)
	return lists, err
}

// SearchLists searches for public lists.
func (c *Client) SearchLists(params url.Values) ([]List, error) {
	var lists []List
	err := c.doRequest(http.MethodGet, "/lists/search", params, nil, &lists)
	return lists, err
}

// GetRatings fetches bulk ratings for media items.
func (c *Client) GetRatings(mediaType, returnRating string, body RatingsRequest) (*RatingsResponse, error) {
	endpoint := fmt.Sprintf("/rating/%s/%s", mediaType, returnRating)
	var ratings RatingsResponse
	err := c.doRequest(http.MethodPost, endpoint, nil, body, &ratings)
	return &ratings, err
}

// ModifyStaticList adds or removes items from a static list.
func (c *Client) ModifyStaticList(listID int, action string, body ModifyListRequest) (*ModifyListResponse, error) {
	endpoint := fmt.Sprintf("/lists/%d/items/%s", listID, action)
	var response ModifyListResponse
	err := c.doRequest(http.MethodPost, endpoint, nil, body, &response)
	return &response, err
}

// GetLastActivities fetches the last activity timestamps.
func (c *Client) GetLastActivities() (*LastActivities, error) {
	var activities LastActivities
	err := c.doRequest(http.MethodGet, "/sync/last_activities", nil, nil, &activities)
	return &activities, err
}

// GetWatchlistItems fetches items from the user's watchlist.
func (c *Client) GetWatchlistItems(params url.Values) (*WatchlistItems, error) {
	var items WatchlistItems
	err := c.doRequest(http.MethodGet, "/watchlist/items", params, nil, &items)
	return &items, err
}

// ModifyWatchlist adds or removes items from the watchlist.
func (c *Client) ModifyWatchlist(action string, body ModifyListRequest) (*ModifyWatchlistResponse, error) {
	endpoint := fmt.Sprintf("/watchlist/items/%s", action)
	var response ModifyWatchlistResponse
	err := c.doRequest(http.MethodPost, endpoint, nil, body, &response)
	return &response, err
}

func (c *Client) doRequest(method, endpoint string, params url.Values, body interface{}, result interface{}) error {
	// Prepare URL
	fullURL, err := url.Parse(apiBaseURL + endpoint)
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}

	query := fullURL.Query()
	if params != nil {
		for key, values := range params {
			for _, value := range values {
				query.Add(key, value)
			}
		}
	}
	query.Set("apikey", c.apiKey)
	fullURL.RawQuery = query.Encode()

	// Prepare request body
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, fullURL.String(), reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal JSON response: %w (body: %s)", err, string(respBody))
		}
	}

	return nil
}
