package endpoints

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	client  http.Client
	baseURL string
	apiKey  string
	token   string
}

func New(projectReference string, apiKey string, isDevelopment bool) *Client {
	var baseURL = "http://127.0.0.1:54321/auth/v1"
	if !isDevelopment {
		baseURL = fmt.Sprintf("https://%s.supabase.co/auth/v1", projectReference)
	}
	return &Client{
		client: http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (c Client) WithCustomAuthURL(url string) *Client {
	return &Client{
		client:  c.client,
		baseURL: url,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}

func (c Client) WithToken(token string) *Client {
	return &Client{
		client:  c.client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   token,
	}
}

func (c Client) WithClient(client http.Client) *Client {
	return &Client{
		client:  client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}

// Returns a copy of a HTTP client that will not follow redirects.
func noRedirClient(client http.Client) http.Client {
	return http.Client{
		Transport: client.Transport,
		Jar:       client.Jar,
		Timeout:   client.Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
