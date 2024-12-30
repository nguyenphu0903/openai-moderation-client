package moderation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const defaultCreateEndpoint = "https://api.openai.com/v1/moderations"

// Client is the moderation API client
type Client struct {
	s              *Session
	model          string
	CreateEndpoint string
}

// Session represents the HTTP session for interacting with the API
type Session struct {
	APIKey string
}

// MakeRequest handles sending HTTP requests
func (s *Session) MakeRequest(ctx context.Context, url string, payload interface{}, result interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.APIKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// NewClient creates a new moderation client
func NewClient(session *Session, model string) *Client {
	return &Client{
		s:              session,
		model:          model,
		CreateEndpoint: defaultCreateEndpoint,
	}
}

// Request represents the moderation request payload
type Request struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

// Response represents the moderation API response
type Response struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Results []Result `json:"results"`
}

// Result represents a single moderation result
type Result struct {
	Categories     map[string]bool    `json:"categories,omitempty"`
	CategoryScores map[string]float64 `json:"category_scores,omitempty"`
	Flagged        bool               `json:"flagged,omitempty"`
}

func (c *Client) Create(ctx context.Context, p *Request) (*Response, error) {
	if p.Model == "" {
		p.Model = c.model
	}

	var r Response
	if err := c.s.MakeRequest(ctx, c.CreateEndpoint, p, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
