package main

type ShipmentState string

const (
	Scheduled ShipmentState = "scheduled"
	Assigned  ShipmentState = "assigned"
	EnRoute   ShipmentState = "en_route"
	Delivered ShipmentState = "delivered"
	Attempted ShipmentState = "attempted"
	Returned  ShipmentState = "returned"
	Canceled  ShipmentState = "canceled"
	Expired   ShipmentState = "expired"
)
