package main

import (
	"time"

	"github.com/kshuta/workoutTracker/data"
)

func getTestWorkouts() []data.Workout {
	workouts := []data.Workout{
		{
			Id:        1,
			Name:      "test-workout",
			WeekNo:    1,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
	}

	return workouts
}

func getTestLifts() []data.Lift {
	lifts := []data.Lift{
		{
			Id:        1,
			Name:      "Bench Press",
			Max:       100,
			CreatedAt: time.Now(),
		},
		{
			Id:        2,
			Name:      "Squat",
			Max:       150,
			CreatedAt: time.Now(),
		},
		{
			Id:        3,
			Name:      "Deadlifts",
			Max:       200,
			CreatedAt: time.Now(),
		},
		{
			Id:        4,
			Name:      "Overhead Press",
			Max:       70,
			CreatedAt: time.Now(),
		},
		{
			Id:        5,
			Name:      "Hang Clean",
			Max:       80,
			CreatedAt: time.Now(),
		},
	}

	return lifts
}

func getTestSets() []data.Set {
	sets := []data.Set{
		{
			Id:        1,
			LiftId:    1,
			Done:      false,
			CreatedAt: time.Now(),
		},
		{
			Id:        2,
			LiftId:    1,
			Done:      false,
			CreatedAt: time.Now(),
		},
		{
			Id:        3,
			LiftId:    1,
			Done:      false,
			CreatedAt: time.Now(),
		},
		{
			Id:        4,
			LiftId:    2,
			Done:      false,
			CreatedAt: time.Now(),
		},
		{
			Id:        5,
			LiftId:    2,
			Done:      false,
			CreatedAt: time.Now(),
		},
		{
			Id:        6,
			LiftId:    3,
			CreatedAt: time.Now(),
		},
		{
			Id:        7,
			LiftId:    3,
			CreatedAt: time.Now(),
		},
	}

	return sets
}

func getTestSetQuantity() []data.SetQuantity {
	sqs := []data.SetQuantity{
		{
			Id:           1,
			SetId:        1,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 60,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
		{
			Id:           2,
			SetId:        2,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 70,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
		{
			Id:           3,
			SetId:        3,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 75,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
		{
			Id:           4,
			SetId:        4,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 65,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
		{
			Id:           5,
			SetId:        5,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 75,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
		{
			Id:           6,
			SetId:        6,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 80,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
		{
			Id:           7,
			SetId:        7,
			Reptype:      data.Count,
			Quantity:     8,
			PlannedRatio: 85,
			Ratiotype:    data.Percentage,
			CreatedAt:    time.Now(),
		},
	}

	return sqs
}
