package client

import (
	"fmt"
	"time"
)

// MyLimits represents the user's API limits.
type MyLimits struct {
	APIRequests      int    `json:"api_requests"`
	APIRequestsCount int    `json:"api_requests_count"`
	UserID           int    `json:"user_id"`
	PatronStatus     string `json:"patron_status"`
	PatreonPledge    int    `json:"patreon_pledge"`
}

// List represents a single list.
type List struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	MediaType   string `json:"mediatype"`
	Items       int    `json:"items"`
	Likes       int    `json:"likes"`
	UserID      int    `json:"user_id,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	Dynamic     bool   `json:"dynamic,omitempty"`
	Private     bool   `json:"private,omitempty"`
}

// ListUpdateResponse represents the response from updating a list.
type ListUpdateResponse struct {
	Success    bool   `json:"success"`
	ID         int    `json:"id,omitempty"`
	UpdatedIDs []int  `json:"updated_ids,omitempty"`
	Name       string `json:"name"`
}

// ListItems represents the items within a list, separated by media type.
type ListItems struct {
	Movies []ListItem `json:"movies"`
	Shows  []ListItem `json:"shows"`
}

// ListItem represents a single movie or show in a list.
type ListItem struct {
	ID             int    `json:"id"`
	Rank           int    `json:"rank"`
	Adult          int    `json:"adult"`
	Title          string `json:"title"`
	ImdbID         string `json:"imdb_id"`
	TvdbID         *int   `json:"tvdb_id"`
	Language       string `json:"language"`
	MediaType      string `json:"mediatype"`
	ReleaseYear    int    `json:"release_year"`
	SpokenLanguage string `json:"spoken_language"`
}

// ListChanges represents the changes in a list.
type ListChanges struct {
	ID    int `json:"id"`
	Movie struct {
		TraktIDs struct {
			Added   []int `json:"added"`
			Removed []int `json:"removed"`
		} `json:"trakt_ids"`
	} `json:"movie"`
	Updated time.Time `json:"updated"`
}

// MediaInfo represents detailed information about a media item.
type MediaInfo struct {
	Title           string `json:"title"`
	Year            int    `json:"year"`
	Released        string `json:"released"`
	ReleasedDigital string `json:"released_digital"`
	Description     string `json:"description"`
	Runtime         int    `json:"runtime"`
	Score           int    `json:"score"`
	ScoreAverage    int    `json:"score_average"`
	IDs             struct {
		Imdb  string `json:"imdb"`
		Trakt int    `json:"trakt"`
		Tmdb  int    `json:"tmdb"`
		Tvdb  *int   `json:"tvdb"`
		Mal   *int   `json:"mal"`
	} `json:"ids"`
	Type    string `json:"type"`
	Ratings []struct {
		Source string      `json:"source"`
		Value  interface{} `json:"value"` // Can be int or float
		Score  *int        `json:"score"`
		Votes  *int        `json:"votes"`
		URL    interface{} `json:"url"` // Can be int or string
	} `json:"ratings"`
	Streams []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"streams"`
	WatchProviders []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"watch_providers"`
	Language       string    `json:"language"`
	SpokenLanguage string    `json:"spoken_language"`
	Country        string    `json:"country"`
	Certification  string    `json:"certification"`
	Commonsense    *bool     `json:"commonsense"`
	AgeRating      *int      `json:"age_rating"`
	Status         string    `json:"status"`
	Trailer        string    `json:"trailer"`
	Poster         string    `json:"poster"`
	Backdrop       string    `json:"backdrop"`
	Reviews        []Review  `json:"reviews,omitempty"`
	Keywords       []Keyword `json:"keywords,omitempty"`
}

// Review represents a media review.
type Review struct {
	UpdatedAt  string `json:"updated_at"`
	Author     string `json:"author"`
	Rating     int    `json:"rating"`
	ProviderID int    `json:"provider_id"`
	Content    string `json:"content"`
}

// Keyword represents a media keyword.
type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MediaInfoBatchRequest represents the request body for a batch media info request.
type MediaInfoBatchRequest struct {
	IDs              []string `json:"ids"`
	AppendToResponse []string `json:"append_to_response,omitempty"`
}

// SearchResult represents the result of a media search.
type SearchResult struct {
	Search []struct {
		Title        string `json:"title"`
		Year         int    `json:"year"`
		Score        int    `json:"score"`
		ScoreAverage int    `json:"score_average"`
		Type         string `json:"type"`
		IDs          struct {
			ImdbID  string `json:"imdbid"`
			TmdbID  int    `json:"tmdbid"`
			TraktID int    `json:"traktid"`
			MalID   *int   `json:"malid"`
			TvdbID  *int   `json:"tvdbid"`
		} `json:"ids"`
	} `json:"search"`
	Total int `json:"total"`
}

// RatingsRequest represents the request body for a bulk ratings request.
type RatingsRequest struct {
	IDs      []int  `json:"ids"`
	Provider string `json:"provider"`
}

// RatingsResponse represents the response from a bulk ratings request.
type RatingsResponse struct {
	ProviderID     string `json:"provider_id"`
	ProviderRating string `json:"provider_rating"`
	MediaType      string `json:"mediatype"`
	Ratings        []struct {
		ID     int     `json:"id"`
		Rating float64 `json:"rating"`
	} `json:"ratings"`
}

// ModifyListRequest represents the request body for adding/removing items from a static list.
type ModifyListRequest struct {
	Movies []map[string]interface{} `json:"movies,omitempty"`
	Shows  []map[string]interface{} `json:"shows,omitempty"`
}

// ModifyListResponse represents the response from modifying a list.
type ModifyListResponse struct {
	Added struct {
		Movies int `json:"movies"`
		Shows  int `json:"shows"`
	} `json:"added"`
	Existing struct {
		Movies int `json:"movies"`
		Shows  int `json:"shows"`
	} `json:"existing"`
	NotFound struct {
		Movies int `json:"movies"`
		Shows  int `json:"shows"`
	} `json:"not_found"`
}

// LastActivities represents the last activity timestamps for sync purposes.
type LastActivities struct {
	WatchlistedAt time.Time `json:"watchlisted_at"`
}

// WatchlistItems represents items in the user's watchlist.
type WatchlistItems struct {
	Movies []WatchlistItem `json:"movies"`
	Shows  []WatchlistItem `json:"shows"`
}

// WatchlistItem represents a single item in the watchlist.
type WatchlistItem struct {
	ListItem
	WatchlistAt string `json:"watchlist_at"`
}

// ModifyWatchlistResponse is an alias for ModifyListResponse as they share the same structure.
type ModifyWatchlistResponse = ModifyListResponse

// APIError represents an error response from the MDBList API.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}
