package main

import (
	"time"

	"github.com/kshuta/workout_tracker/data"
)

var workout data.Workout = data.Workout{
	Id:        1,
	Name:      "test-workout",
	WeekNo:    1,
	Date:      time.Now(),
	CreatedAt: time.Now(),
}
