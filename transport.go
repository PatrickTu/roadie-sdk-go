package main

import (
	"encoding/base64"
	"net/http"
)

// AuthBasicTransport supports basic authentication using the authorization header
// https://docs.roadie.com/#authentication-methods
type AuthBasicTransport struct {
	username string `json:"-"`
	password string `json:"-"`
}

func (t *AuthBasicTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	auth := t.username + ":" + t.password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Add("Authorization", "Bearer "+encoded)
	return http.DefaultTransport.RoundTrip(req)
}

func NewAuthBasicTransport(username string, password string) *AuthBasicTransport {
	return &AuthBasicTransport{username: username, password: password}
}

// AuthAPIKeyTransport supports bearer tokens using the authorization header
// https://docs.roadie.com/#authentication-methods
type AuthBearerTokenTransport struct {
	bearerToken string `json:"-"`
}

func (t *AuthBearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+t.bearerToken)
	return http.DefaultTransport.RoundTrip(req)
}

func NewAuthBearerTokenTransport(bearerToken string) *AuthBearerTokenTransport {
	return &AuthBearerTokenTransport{bearerToken: bearerToken}
}

// AuthAPIKeyTransport supports API keys using the x-api-key header
// https://docs.roadie.com/#authentication-methods
type AuthAPIKeyTransport struct {
	apiKey string `json:"-"`
}

func (t *AuthAPIKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-api-key", t.apiKey)
	return http.DefaultTransport.RoundTrip(req)
}

func NewAuthAPIKeyTransport(apiKey string) *AuthAPIKeyTransport {
	return &AuthAPIKeyTransport{apiKey: apiKey}
}
