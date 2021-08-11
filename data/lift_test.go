package data

import (
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

	t.Run("create lift with empty CreatedAt (should fail)", func(t *testing.T) {
		t.Parallel()
		lift := getTestLift("create test lift name")
		lift.CreatedAt = time.Time{}

		err := lift.Create()
		assertError(t, err, ErrLiftMissingField)

	})
}

func liftIsCreated(t *testing.T, lift Lift, err error) {
	assertNoError(t, err)

	if lift.Id == 0 {
		t.Error("insertion failed: lift id is still 0")
	}
}

func TestLiftRetrieve(t *testing.T) {
	t.Parallel()
	t.Run("retrieve lift", func(t *testing.T) {
		lift := getTestLift("retrieve test lift name")
		err := lift.Create()
		liftIsCreated(t, *lift, err)

		retrievedLift, err := GetLift(lift.Id)
		assertNoError(t, err)
		if retrievedLift.Id != lift.Id {
			t.Errorf("Expected lift with id %d, got lift with id %d", lift.Id, retrievedLift.Id)
		}
	})

	t.Run("retrieve lift that doesn't exist", func(t *testing.T) {
		_, err := GetLift(-1)
		assertError(t, err, err)
	})

}

func getTestLift(liftName string) (lift *Lift) {
	lift = &Lift{
		Name:      liftName,
		Max:       60,
		CreatedAt: time.Now(),
	}
	return
}
