package main

import "time"

type ShipmentService interface {
	CreateShipment(CreateShipmentRequest) (*Shipment, error)
	RetrieveShipment(id int) (*Shipment, error)
	RetrieveShipments(ids []int, referenceIds []string) ([]*Shipment, error)
	UpdateShipment(UpdateShipmentRequest) (*Shipment, error)
	CancelShipment(CancelShipmentRequest) error // 200 or 204 for no content
}

type Shipment struct {
	ID                int             `json:"id"`
	ReferenceID       string          `json:"reference_id"`
	Description       string          `json:"description"`
	State             string          `json:"state"`
	AlternateID1      string          `json:"alternate_id_1"`
	AlternateID2      string          `json:"alternate_id_2"`
	Items             []Item          `json:"items"`
	PickupLocation    Location        `json:"pickup_location"`
	DeliveryLocation  Location        `json:"delivery_location"`
	PickupAfter       time.Time       `json:"pickup_after"`
	DeliverBetween    TimeWindow      `json:"deliver_between"`
	Options           DeliveryOptions `json:"options"`
	TrackingNumber    string          `json:"tracking_number"`
	Price             float64         `json:"price"`
	EstimatedDistance float64         `json:"estimated_distance"`
	Events            []Event         `json:"events"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

type Event struct {
	Name       string        `json:"name"`
	OccurredAt time.Time     `json:"occurred_at"`
	Location   EventLocation `json:"location"`
}

type EventLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateShipmentRequest struct {
	ReferenceID      string          `json:"reference_id"` // required
	IdempotencyKey   string          `json:"idempotency_key"`
	AlternateID1     string          `json:"alternate_id_1"`
	AlternateID2     string          `json:"alternate_id_2"`
	Description      string          `json:"description"`
	Items            []Item          `json:"items"`             // required
	PickupLocation   Location        `json:"pickup_location"`   // required
	DeliveryLocation Location        `json:"delivery_location"` // required
	PickupAfter      time.Time       `json:"pickup_after"`      // required
	PickupBetween    TimeWindow      `json:"pickup_between"`
	DeliverBetween   TimeWindow      `json:"deliver_between"` // required
	Options          DeliveryOptions `json:"options"`         // required
	Events           []Event         `json:"events"`
}

type UpdateShipmentRequest struct {
	ReferenceID      string          `json:"reference_id"`
	AlternateID1     string          `json:"alternate_id_1"`
	AlternateID2     string          `json:"alternate_id_2"`
	Description      string          `json:"description"`
	Items            []Item          `json:"items"`
	PickupLocation   Location        `json:"pickup_location"`
	DeliveryLocation Location        `json:"delivery_location"`
	PickupAfter      time.Time       `json:"pickup_after"`
	DeliverBetween   TimeWindow      `json:"deliver_between"`
	Options          DeliveryOptions `json:"options"`
}

type CancelShipmentRequest struct {
	CancellationCode    string `json:"cancellation_code"`
	CancellationComment string `json:"cancellation_comment"`
}
