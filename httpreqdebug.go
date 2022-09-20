package util

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// PrettyHTTPRequest ...
type PrettyHTTPRequest struct {
	Host   string      `json:"host,omitempty"`
	Method string      `json:"method,omitempty"`
	URL    *url.URL    `json:"url,omitempty"`
	Header http.Header `json:"header,omitempty"`
}

// NewPrettyHTTPRequest ...
func NewPrettyHTTPRequest(r *http.Request) *PrettyHTTPRequest {

	return &PrettyHTTPRequest{
		Host:   r.Host,
		Method: r.Method,
		URL:    r.URL,
		Header: r.Header,
	}

}

func (t *PrettyHTTPRequest) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}
