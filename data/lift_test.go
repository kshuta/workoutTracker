package data

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestLiftCreate(t *testing.T) {
	t.Parallel()

	t.Run("create lift", func(t *testing.T) {
		t.Parallel()
		lift := getTestLift("create test lift name")

		err := lift.Create()
		liftIsCreated(t, *lift, err)
	})

	t.Run("create lift with empty max (should succeed)", func(t *testing.T) {
		t.Parallel()
		lift := getTestLift("create test lift name")

		lift.Max = 0
		err := lift.Create()
		liftIsCreated(t, *lift, err)
	})
}

func liftIsCreated(t *testing.T, lift Lift, err error) {
	t.Helper()
	assertNoError(t, err)

	if lift.Id == 0 {
		t.Error("insertion failed: lift id is still 0")
	}
}

func TestLiftRetrieve(t *testing.T) {
	t.Parallel()
	t.Run("retrieve lift", func(t *testing.T) {
		t.Parallel()
		lift := getTestLift("retrieve test lift name")
		err := lift.Create()
		assertNoError(t, err)
		liftIsCreated(t, *lift, err)

		retrievedLift, err := GetLift(lift.Id)
		assertNoError(t, err)
		if retrievedLift.Id != lift.Id {
			t.Errorf("Expected lift with id %d, got lift with id %d", lift.Id, retrievedLift.Id)
		}
	})

	// must implement (but I don't know how right now)
	t.Run("retrieve all lifts", func(t *testing.T) {

	})

	t.Run("retrieve lift that doesn't exist", func(t *testing.T) {
		t.Parallel()
		_, err := GetLift(-1)
		assertError(t, err, sql.ErrNoRows)
	})

}

func TestLiftUpdate(t *testing.T) {
	t.Parallel()
	beforeUpdate := "not updated"
	lift := getTestLift(beforeUpdate)
	err := lift.Create()
	liftIsCreated(t, *lift, err)

	afterUpdate := "updated"
	lift.Name = afterUpdate
	err = lift.Update()

	assertNoError(t, err)

	retrievedLift, err := GetLift(lift.Id)
	assertNoError(t, err)

	if retrievedLift.Name != afterUpdate {
		t.Error("error: name not updated")
	}
}

func TestLiftDelete(t *testing.T) {
	t.Parallel()
	delLift := getTestLift("lift soon to be deleted")
	err := delLift.Create()
	liftIsCreated(t, *delLift, err)

	err = delLift.Delete()
	assertNoError(t, err)

	_, err = GetLift(delLift.Id)
	assertError(t, err, sql.ErrNoRows)
}

func TestLiftWorkout(t *testing.T) {
	liftNames := []string{
		"Benchpress",
		"Squat",
		"Deadlift",
		"OverHead press",
	}

	lifts := make([]Lift, len(liftNames))
	workouts := make([]Workout, 3)

	for idx, val := range liftNames {
		lift := getTestLift(val)
		err := lift.Create()
		liftIsCreated(t, *lift, err)
		lifts[idx] = *lift
	}

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("Workout:%d", i+1)
		workout := getTestWorkout(name)
		workout.WeekNo = i + 1
		err := workout.Create()
		workoutIsCreated(t, *workout, err)
		workouts[i] = *workout
	}
	t.Run("create TestLiftWorkout", func(t *testing.T) {
		for idx, workout := range workouts {
			for _, lift := range lifts[idx : idx+2] {
				err := CreateLiftWorkout(&workout, &lift)
				assertNoError(t, err)
			}
		}
	})

	t.Run("retrieve lifts associated with workout", func(t *testing.T) {
		expectedNames := [][]string{
			{
				"Benchpress",
				"Squat",
			},
			{
				"Squat",
				"Deadlift",
			},
			{
				"Deadlift",
				"OverHead press",
			},
		}

		for widx, names := range expectedNames {
			lifts, err := GetWorkoutLifts(workouts[widx])
			assertNoError(t, err)

			found := false
			for _, name := range names {
				for _, lift := range lifts {
					if lift.Name == name {
						found = true
						break
					}
				}

				if !found {
					t.Fatalf("could not find expected lift: %s", name)
				}
			}
		}
	})

	t.Run("get all lifts", func(t *testing.T) {
		_, err := db.Exec("update lifts set is_deleted = true") // reset lift table
		if err != nil {
			t.Error(err)
		}

		for i := 0; i < 4; i++ {
			lift := getTestLift("get all lift test")
			err := lift.Create()
			assertNoError(t, err)
		}

		lifts, err := GetLifts()
		assertNoError(t, err)

		if len(lifts) != 4 {
			t.Fatalf("could not retrieve all lifts")
		}

	})
}

// returns lift struct with populated fields
func getTestLift(liftName string) (lift *Lift) {
	lift = &Lift{
		Name:      liftName,
		Max:       60,
		CreatedAt: time.Now(),
		IsDeleted: false,
	}
	return
}
