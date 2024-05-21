package test

import (
	"net/http"
)

// RoundTripFunc ...
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip ...
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestHTTPClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestHTTPClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
