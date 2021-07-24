package data

import "time"

type Plans struct {
	Id        int
	name      string
	duration  int
	frequency int
}

type Workouts struct {
	Id     int
	planId int
	name   string
	weekNo int
	date   time.Time
}

type Lifts struct {
	Id        int
	workoutId int
	name      string
	max       float64
}

type Sets struct {
	Id     int
	liftId int
	done   bool
}

type SetQuantities struct {
	Id           int
	SetId        int
	repType      string
	quantity     int
	weight       float64
	plannedRatio int
	ratiotype    string
}
