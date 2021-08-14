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
		setIsCreated(t, *set, err)
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

func TestSetRetrieve(t *testing.T) {
	t.Parallel()
	t.Run("retrieve set", func(t *testing.T) {
		set, err := getTestSet("test set name")
		assertNoError(t, err)
		err = set.Create()
		setIsCreated(t, *set, err)

		retrievedSet, err := GetSet(set.Id)
		assertNoError(t, err)

		if retrievedSet.Id != set.Id {
			t.Errorf("Expected to retrieve set with id %d, got set with id %d", set.Id, retrievedSet.Id)
		}
	})

	t.Run("retrieve set that doesn't exist", func(t *testing.T) {
		_, err := GetSet(-1)
		assertError(t, err, err)
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

func setIsCreated(t *testing.T, set Set, err error) {
	assertNoError(t, err)
	if set.Id == 0 {
		t.Error("Insersion failed: lift id is still 0")
	}
}
