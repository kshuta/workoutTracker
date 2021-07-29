package data

import (
	"testing"
	"time"
)

func TestWorkoutCreate(t *testing.T) {
	t.Parallel()

	plan := Plan{
		Name:      "Workout test plan Name",
		Duration:  4,
		Frequency: 3,
		CreatedAt: time.Now(),
	}

	err := plan.Create()
	assertNoError(t, err)

	t.Run("creating workout", func(t *testing.T) {
		workout := Workout{
			Name:      "Test workout",
			WeekNo:    1,
			Date:      time.Now(),
			CreatedAt: time.Now(),
			PlanId:    plan.Id,
		}

		err = workout.Create()
		assertNoError(t, err)
		if workout.Id == 0 {
			t.Error("insertion failed: workout id is still 0")
		}

	})

	t.Run("creating workout without PlanId", func(t *testing.T) {
		workout := Workout{
			Name:      "Test workout",
			WeekNo:    1,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		}
		err = workout.Create()
		testWorkoutEmptyField(t, workout, err)
	})

	t.Run("creating workout without Name", func(t *testing.T) {
		workout := Workout{
			PlanId:    plan.Id,
			WeekNo:    1,
			Date:      time.Now(),
			CreatedAt: time.Now(),
		}
		err = workout.Create()
		testWorkoutEmptyField(t, workout, err)
	})

	t.Run("creating workout without WeekNo", func(t *testing.T) {
		workout := Workout{
			PlanId:    plan.Id,
			Name:      "Test workout",
			Date:      time.Now(),
			CreatedAt: time.Now(),
		}
		err = workout.Create()
		testWorkoutEmptyField(t, workout, err)
	})

	t.Run("creating workout without Date", func(t *testing.T) {
		workout := Workout{
			PlanId:    plan.Id,
			Name:      "Test workout",
			WeekNo:    1,
			CreatedAt: time.Now(),
		}
		err = workout.Create()
		testWorkoutEmptyField(t, workout, err)
	})

	t.Run("creating workout without CreatedAt", func(t *testing.T) {
		workout := Workout{
			PlanId: plan.Id,
			Name:   "Test workout",
			WeekNo: 1,
			Date:   time.Now(),
		}
		err = workout.Create()
		testWorkoutEmptyField(t, workout, err)
	})
}

// checks for correct error const
// and also if insertion failed
func testWorkoutEmptyField(t *testing.T, workout Workout, err error) {
	t.Helper()
	assertError(t, err, ErrWorkoutMissingField)
	if workout.Id != 0 {
		t.Error("error: insertion did not fail with empty field(s)")
	}
}

func TestWorkoutRetrieve(t *testing.T) {
	t.Parallel()
	plan := Plan{
		Name:      "Workout retrieve test plan Name",
		Duration:  4,
		Frequency: 3,
		CreatedAt: time.Now(),
	}
	err := plan.Create()
	assertNoError(t, err)

	workout := Workout{
		Name:      "Test retrieve workout",
		WeekNo:    1,
		Date:      time.Now(),
		CreatedAt: time.Now(),
		PlanId:    plan.Id,
	}
	err = workout.Create()
	assertNoError(t, err)

	retreivedWorkout, err := GetWorkout(workout.Id)
	assertNoError(t, err)
	if retreivedWorkout.Id != workout.Id {
		t.Errorf("wrong workout retreived: wanted workout with id %d, but retreived workout with id %d", workout.Id, retreivedWorkout.Id)
	}

}
