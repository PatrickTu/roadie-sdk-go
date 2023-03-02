package main

import "time"

type Item struct {
	Description string  `json:"description"` // required
	ReferenceId string  `json:"reference_id"`
	Value       float64 `json:"value"`    // required
	Length      float64 `json:"length"`   // required
	Width       float64 `json:"width"`    // required
	Height      float64 `json:"height"`   // required
	Weight      float64 `json:"weight"`   // required
	Quantity    int     `json:"quantity"` // required
}

type Location struct {
	Address Address `json:"address"` // required
	Contact Contact `json:"contact"` // required for shipments
	Notes   string  `json:"notes"`   // max 500 characters
}

type Contact struct {
	Name  string `json:"name"`  // required
	Phone string `json:"phone"` // required
}

type Address struct {
	Name        string  `json:"name"`         // max 200 characters
	StoreNumber string  `json:"store_number"` // max 20 characters
	Street1     string  `json:"street1"`      // required, max 200 characters
	Street2     string  `json:"street2"`      // max 200 characters
	City        string  `json:"city"`         // required, max 200 characters
	State       string  `json:"state"`        // required
	Zip         string  `json:"zip"`          // required, max 10 characters
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type DeliveryOptions struct {
	SignatureRequired    bool    `json:"signature_required"` // required
	NotificationsEnabled bool    `json:"notifications_enabled"`
	Over21Required       bool    `json:"over_21_required"`
	ExtraCompensation    float64 `json:"extra_compensation"`
	TrailerRequired      bool    `json:"trailer_required"`
	DeclineInsurance     bool    `json:"decline_insurance"`
}

// TimeWindow window must be 2 hours or greater
type TimeWindow struct {
	Start time.Time `json:"start"` // required
	End   time.Time `json:"end"`   // required
}

type Driver struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
