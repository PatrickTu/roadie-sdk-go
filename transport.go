package main

import (
	"encoding/base64"
	"net/http"
)

// AuthBasicTransport supports basic authentication using the authorization header
// https://docs.roadie.com/#authentication-methods
type AuthBasicTransport struct {
	Username string
	Password string `json:"-"`
}

func (t *AuthBasicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	auth := t.Username + ":" + t.Password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	
	req.Header.Add("Authorization", "Bearer "+encoded)
	return http.DefaultTransport.RoundTrip(req)
}

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
