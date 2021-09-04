package main

import (
	"fmt"
	"time"

	"github.com/kshuta/workoutTracker/data"
)

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
