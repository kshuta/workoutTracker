package data

import (
	"testing"
	"time"
)

func TestSetCreate(t *testing.T) {
	t.Parallel()
	t.Run("create set", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet("Test set name")
		assertNoError(t, err)
		err = set.Create()
		assertNoError(t, err)
		if set.Id == 0 {
			t.Error("Insersion failed: lift id is still 0")
		}
	})

	t.Run("create set without lift_id", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet("test set name")
		assertNoError(t, err)
		set.LiftId = 0
		err = set.Create()
		assertError(t, err, ErrSetMissingField)
	})

	t.Run("create set without created_at", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet("test set name")
		assertNoError(t, err)
		set.CreatedAt = time.Time{}
		err = set.Create()
		assertError(t, err, ErrSetMissingField)
	})
}

// returns set struct with populated fields
// creates arbitrary lift for parent
func getTestSet(setName string) (set *Set, err error) {
	lift := getTestLift("lift for set test")
	err = lift.Create()
	if err != nil {
		return nil, err
	}

	set = &Set{
		LiftId:    lift.Id,
		Done:      false,
		CreatedAt: time.Now(),
	}

	return
}
