package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RoadieClient struct {
	// Represents the base url used to interact with the API
	baseUrl string

	// HTTP client that will be used to perform requests
	client *http.Client
}

// NewRoadieClient serves as the main constructor to create a Roadie client
func NewRoadieClient(baseUrl string, client *http.Client) *RoadieClient {
	return &RoadieClient{baseUrl: baseUrl, client: client}
}

func (c *RoadieClient) prepareBody(request interface{}) (*bytes.Buffer, error) {
	encoded, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(encoded), nil
}

// CreateEstimate
// https://docs.roadie.com/#create-an-estimate
func (c *RoadieClient) CreateEstimate(request CreateEstimateRequest) (*Estimate, error) {
	body, err := c.prepareBody(request)
	if err != nil {
		return nil, err
	}

	// Perform request
	response, err := c.client.Post(c.baseUrl+"/v1/estimates", "application/json", body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Parse response
	var estimate Estimate
	if err := json.NewDecoder(response.Body).Decode(&estimate); err != nil {
		return nil, err
	}

	return &estimate, nil
}
