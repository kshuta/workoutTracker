package data

import "time"

type Lift struct {
	Id        int
	WorkoutId int
	Name      string
	Max       float64
	CreatedAt time.Time
}
