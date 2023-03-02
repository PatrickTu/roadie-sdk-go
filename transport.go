package main

import "net/http"

// AuthAPIKeyTransport supports bearer tokens using the authorization header
// https://docs.roadie.com/#authentication-methods
type AuthBearerTokenTransport struct {
	BearerToken string
}

func (t *AuthBearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.BearerToken)
	return http.DefaultTransport.RoundTrip(req)
}

// AuthAPIKeyTransport supports API keys using the x-api-key header
// https://docs.roadie.com/#authentication-methods
type AuthAPIKeyTransport struct {
	APIKey string
}

func (t *AuthAPIKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-api-key", t.APIKey)
	return http.DefaultTransport.RoundTrip(req)
}
