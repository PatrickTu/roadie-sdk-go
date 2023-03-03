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

func TestRoadieClient_CreateShipment(t *testing.T) {
	client := GetTestRoadieClient()
	request := CreateShipmentRequest{Items: []Item{}}

	shipment, err := client.CreateShipment(request)

	if err != nil {
		t.Error(err)
	}

	if shipment.ID != 152040 || shipment.ReferenceID != "ABCDEFG12345" || shipment.State != "scheduled" {
		t.Error("object has wrong values")
	}
}

func TestRoadieClient_RetrieveShipment(t *testing.T) {
	client := GetTestRoadieClient()
	shipment, err := client.RetrieveShipment(152040)

	if err != nil {
		t.Error(err)
	}

	if shipment.ID != 152040 || shipment.ReferenceID != "ABCDEFG12345" || shipment.State != "delivered" {
		t.Error("object has wrong values")
	}
}

func TestRoadieClient_RetrieveShipments(t *testing.T) {
	client := GetTestRoadieClient()
	shipments, err := client.RetrieveShipments([]int{}, []string{"ABC123", "ABC456"})
	if err != nil {
		t.Error(err)
	}

	if len(shipments) != 2 {
		t.Error("shipments does not have 2 items")
	}
}

func TestRoadieClient_UpdateShipment(t *testing.T) {
	client := GetTestRoadieClient()
	shipment, err := client.UpdateShipment(152040, UpdateShipmentRequest{})
	if err != nil {
		t.Error(err)
	}

	if shipment.ID != 152040 || shipment.ReferenceID != "ABCDEFG12345" || shipment.State != "scheduled" {
		t.Error("object has wrong values")
	}
}

func TestRoadieClient_CancelShipment(t *testing.T) {
	client := GetTestRoadieClient()
	err := client.CancelShipment(152040, CancelShipmentRequest{})
	if err != nil {
		t.Error(err)
	}
}
