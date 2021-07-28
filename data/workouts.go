package data

import "time"

type Plan struct {
	Id        int
	Name      string
	Duration  int
	Frequency int
	CreatedAt time.Time `db:"created_at"`
}

type Workout struct {
	Id        int
	PlanId    int
	Name      string
	WeekNo    int
	Date      time.Time
	CreatedAt time.Time
}

type Lift struct {
	Id        int
	WorkoutId int
	Name      string
	Max       float64
	CreatedAt time.Time
}

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
