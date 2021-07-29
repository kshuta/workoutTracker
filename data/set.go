package data

import "time"

type Set struct {
	Id        int
	LiftId    int
	Done      bool
	CreatedAt time.Time
}

type SetQuantity struct {
	Id           int
	SetId        int
	RepType      string
	Quantity     int
	Weight       float64
	PlannedRatio int
	Ratiotype    string
	CreatedAt    time.Time
}
