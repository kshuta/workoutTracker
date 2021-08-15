package data

import (
	"testing"
	"time"
)

func TestWorkoutCreate(t *testing.T) {
	t.Parallel()

	t.Run("creating workout", func(t *testing.T) {
		workout := getTestWorkout("Test workout")

		err := workout.Create()
		assertNoError(t, err)
		if workout.Id == 0 {
			t.Error("insertion failed: workout id is still 0")
		}

	})

	t.Run("creating workout without Name", func(t *testing.T) {
		workout := getTestWorkout("Test workout")
		workout.Name = ""
		err := workout.Create()
		testWorkoutEmptyField(t, *workout, err)
	})

	t.Run("creating workout without WeekNo", func(t *testing.T) {
		workout := getTestWorkout("Test workout")
		workout.WeekNo = 0
		err := workout.Create()
		testWorkoutEmptyField(t, *workout, err)
	})

	t.Run("creating workout without Date", func(t *testing.T) {
		workout := getTestWorkout("Test workout")
		workout.Date = time.Time{}
		err := workout.Create()
		testWorkoutEmptyField(t, *workout, err)
	})

	t.Run("creating workout without CreatedAt", func(t *testing.T) {
		workout := getTestWorkout("Test workout")
		workout.CreatedAt = time.Time{}
		err := workout.Create()
		testWorkoutEmptyField(t, *workout, err)
	})
}

// checks for correct error const
// and also if insertion fails
func testWorkoutEmptyField(t *testing.T, workout Workout, err error) {
	t.Helper()
	assertError(t, err, ErrWorkoutMissingField)
	if workout.Id != 0 {
		t.Error("error: insertion did not fail with empty field(s)")
	}
}

func TestWorkoutRetrieve(t *testing.T) {
	t.Parallel()

	t.Run("retrieving workout", func(t *testing.T) {
		t.Parallel()
		workout := getTestWorkout("workout retrieve test")

		err := workout.Create()
		assertNoError(t, err)

		retreivedWorkout, err := GetWorkout(workout.Id)
		assertNoError(t, err)
		if retreivedWorkout.Id != workout.Id {
			t.Errorf("wrong workout retreived: wanted workout with id %d, but retreived workout with id %d", workout.Id, retreivedWorkout.Id)
		}
	})

	t.Run("retrieving workout that doesn't exist", func(t *testing.T) {
		t.Parallel()
		_, err := GetWorkout(10000000)
		assertError(t, err, err)
	})
}

func TestWorkoutUpdate(t *testing.T) {
	t.Parallel()
	t.Run("updating field", func(t *testing.T) {
		workout := getTestWorkout("before update workout name")
		err := workout.Create()
		assertNoError(t, err)

		updatedName := "updated workout name"
		workout.Name = updatedName
		err = workout.Update()
		assertNoError(t, err)

		updatedWorkout, err := GetWorkout(workout.Id)
		assertNoError(t, err)
		if updatedWorkout.Name != updatedName {
			t.Errorf("Wanted workout name to be updated to '%s', but was still '%s'", updatedName, updatedWorkout.Name)
		}
	})
}

func TestWorkoutDelete(t *testing.T) {
	t.Parallel()
	t.Run("deleting workout", func(t *testing.T) {
		workout := getTestWorkout("to be deleted workout name")
		err := workout.Create()
		assertNoError(t, err)

		workout.Delete()

		_, err = GetWorkout(workout.Id)
		assertError(t, err, err)
	})
}

// returns workout struct with populated fields
func getTestWorkout(workoutName string) (workout *Workout) {
	workout = &Workout{
		Name:      workoutName,
		WeekNo:    1,
		Date:      time.Now(),
		CreatedAt: time.Now(),
	}
	return
}
