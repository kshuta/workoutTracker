package data

import "testing"

func LiftWorkoutTest(t *testing.T) {
	workout := getTestWorkout("lift workout test")
	err := workout.Create()
	if err != nil {
		t.Fatal(err)
	}

	lift := getTestLift("lift workout test")
	err = lift.Create()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Create LiftWorkout", func(t *testing.T) {
		// there's no functionality to retrieve LiftWorkout for now, so the test passes if there is no error returned from LiftWorkout
		err = CreateLiftWorkout(workout, lift)
		assertNoError(t, err)
	})
	t.Run("get lifts related to workout", func(t *testing.T) {
		lifts, err := GetWorkoutLifts(*workout)
		assertNoError(t, err)
		if lifts[0].Id != lift.Id {
			t.Errorf("Expected lift with id %d, got lift with id %d", lift.Id, lifts[0].Id)
		}
	})
}
