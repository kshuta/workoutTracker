package main

import (
	"fmt"
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
		{
			Id:        2,
			Name:      "test-workout-2",
			WeekNo:    2,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
		{
			Id:        3,
			Name:      "test-workout-3",
			WeekNo:    2,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
		{
			Id:        4,
			Name:      "test-workout-4",
			WeekNo:    2,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
		{
			Id:        5,
			Name:      "test-workout-5",
			WeekNo:    2,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
		{
			Id:        6,
			Name:      "test-workout-6",
			WeekNo:    2,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
		{
			Id:        7,
			Name:      "test-workout-7",
			WeekNo:    2,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		},
		{
			Id:        8,
			Name:      "test-workout-8",
			WeekNo:    2,
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

type LiftInfo struct {
	Lift     data.Lift
	Setinfos []SetInfo
}
type SetInfo struct {
	Set      data.Set
	Quantity data.SetQuantity
}

type WorkoutInfo struct {
	Workout   data.Workout
	Liftinfos []LiftInfo
}

func createTestData() (workoutinfos []WorkoutInfo, err error) {

	for i := 0; i < 8; i++ {
		name := fmt.Sprintf("workout%d", i)
		weekno := i%4 + 1
		workout := data.Workout{
			Name:      name,
			WeekNo:    weekno,
			Date:      time.Date(2021, time.August, 30+i, 0, 0, 0, 0, time.Local),
			CreatedAt: time.Now(),
		}
		err = workout.Create()
		if err != nil {
			return
		}

		liftinfos := make([]LiftInfo, 0)
		for li := 0; li < 2; li++ {
			name := fmt.Sprintf("lift%d", (li+i)%4+1)
			max := float64(50 * ((li+i)%4 + 1))
			lift := data.Lift{
				Name:      name,
				Max:       max,
				CreatedAt: time.Now(),
			}

			err = lift.Create()

			if err != nil {
				return
			}

			setinfos := make([]SetInfo, 0)
			for si := 0; si < 4; si++ {
				set := data.Set{
					LiftId:    lift.Id,
					WorkoutId: workout.Id,
					Done:      false,
					CreatedAt: time.Now(),
				}
				err = set.Create()
				if err != nil {
					return
				}

				sq := data.SetQuantity{
					SetId:        set.Id,
					Reptype:      data.Count,
					Quantity:     i + 1,
					PlannedRatio: int(70 + (si * 3)),
					Ratiotype:    data.Percentage,
					CreatedAt:    time.Now(),
				}

				sq.Weight = calcWeight(lift, sq)

				setinfo := SetInfo{
					Set:      set,
					Quantity: sq,
				}
				setinfos = append(setinfos, setinfo)
			}
			liftinfos = append(liftinfos, LiftInfo{
				Lift:     lift,
				Setinfos: setinfos,
			})
		}

		workoutinfos = append(workoutinfos, WorkoutInfo{
			Workout:   workout,
			Liftinfos: liftinfos,
		})
	}

	return

}
