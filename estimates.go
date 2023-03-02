package main

import "time"

type EstimateService interface {
	CreateEstimate(CreateEstimateRequest) (*Estimate, error)
}

type CreateEstimateRequest struct {
	Items            []Item     `json:"items"`
	PickupLocation   Location   `json:"pickup_location"`
	DeliveryLocation Location   `json:"delivery_location"`
	PickupAfter      time.Time  `json:"pickup_after"`
	DeliverBetween   TimeWindow `json:"deliver_between"`
}

type Estimate struct {
	Price             float64 `json:"price"`
	Size              string  `json:"size"`
	EstimatedDistance float64 `json:"estimated_distance"`
}
