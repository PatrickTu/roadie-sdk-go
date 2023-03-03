package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func (c *RoadieClient) prepareRequest(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Always uses JSON
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}

// CreateEstimate
// https://docs.roadie.com/#create-an-estimate
func (c *RoadieClient) CreateEstimate(request CreateEstimateRequest) (Estimate, error) {
	body, err := c.prepareBody(request)
	if err != nil {
		return Estimate{}, err
	}

	// Perform request
	response, err := c.client.Post(c.baseUrl+"/v1/estimates", "application/json", body)
	if err != nil {
		return Estimate{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Estimate{}, fmt.Errorf("got status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Parse response
	var estimate Estimate
	if err := json.NewDecoder(response.Body).Decode(&estimate); err != nil {
		return Estimate{}, err
	}

	return estimate, nil
}

func (c *RoadieClient) CreateShipment(request CreateShipmentRequest) (Shipment, error) {
	body, err := c.prepareBody(request)
	if err != nil {
		return Shipment{}, err
	}

	req, err := c.prepareRequest("POST", fmt.Sprintf("%s/v1/shipments", c.baseUrl), body)
	if err != nil {
		return Shipment{}, err
	}

	// Perform request
	response, err := c.client.Do(req)
	if err != nil {
		return Shipment{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Shipment{}, fmt.Errorf("got status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Parse response
	var shipment Shipment
	if err := json.NewDecoder(response.Body).Decode(&shipment); err != nil {
		return Shipment{}, err
	}

	return shipment, nil
}

func (c *RoadieClient) RetrieveShipment(id int) (Shipment, error) {
	req, err := c.prepareRequest("GET", fmt.Sprintf("%s/v1/shipments/%d", c.baseUrl, id), nil)
	if err != nil {
		return Shipment{}, err
	}

	// Perform request
	response, err := c.client.Do(req)
	if err != nil {
		return Shipment{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Shipment{}, fmt.Errorf("got status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Parse response
	var shipment Shipment
	if err := json.NewDecoder(response.Body).Decode(&shipment); err != nil {
		return Shipment{}, err
	}

	return shipment, nil
}

func (c *RoadieClient) RetrieveShipments(ids []int, referenceIds []string) ([]Shipment, error) {
	req, err := c.prepareRequest("GET", fmt.Sprintf("%s/v1/shipments", c.baseUrl), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if len(ids) > 0 {
		q.Add("ids", arrayToString(ids, ","))

	}
	if len(referenceIds) > 0 {
		q.Add("reference_ids", strings.Join(referenceIds[:], ","))
	}

	req.URL.RawQuery = q.Encode()

	// Perform request
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Parse response
	var shipments []Shipment

	if err := json.NewDecoder(response.Body).Decode(&shipments); err != nil {
		return nil, err
	}

	return shipments, nil
}

func (c *RoadieClient) UpdateShipment(id int, request UpdateShipmentRequest) (Shipment, error) {
	body, err := c.prepareBody(request)
	if err != nil {
		return Shipment{}, err
	}

	req, err := c.prepareRequest("PATCH", fmt.Sprintf("%s/v1/shipments/%d", c.baseUrl, id), body)
	if err != nil {
		return Shipment{}, err
	}

	// Perform request
	response, err := c.client.Do(req)
	if err != nil {
		return Shipment{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Shipment{}, fmt.Errorf("got status code %d", response.StatusCode)
	}

	defer response.Body.Close()

	// Parse response
	var shipment Shipment
	if err := json.NewDecoder(response.Body).Decode(&shipment); err != nil {
		return Shipment{}, err
	}

	return shipment, nil
}

func (c *RoadieClient) CancelShipment(id int, request CancelShipmentRequest) error {
	body, err := c.prepareBody(request)
	if err != nil {
		return err
	}

	req, err := c.prepareRequest("DELETE", fmt.Sprintf("%s/v1/shipments/%d", c.baseUrl, id), body)
	if err != nil {
		return err
	}

	// Perform request
	response, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("got status code %d", response.StatusCode)
	}

	response.Body.Close()

	return nil
}
