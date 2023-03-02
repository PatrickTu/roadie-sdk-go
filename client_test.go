package main

import (
	"net/http"
	"testing"
)

// Configured to work with Mockoon Mock API
func GetTestRoadieClient() *RoadieClient {
	return NewRoadieClient("http://localhost:9080", &http.Client{Transport: NewAuthAPIKeyTransport("key")})
}

func TestRoadieClient_CreateEstimate(t *testing.T) {
	client := GetTestRoadieClient()
	request := CreateEstimateRequest{Items: []Item{}}

	estimate, err := client.CreateEstimate(request)

	if err != nil {
		t.Error(err)
	}

	if estimate.Price != 12 || estimate.Size != "large" || estimate.EstimatedDistance != 12.54 {
		t.Error("object has wrong values")
	}
}
