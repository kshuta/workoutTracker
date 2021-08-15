package data

import (
	"testing"
	"time"
)

func TestSetCreate(t *testing.T) {
	t.Parallel()
	t.Run("create set", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		err = set.Create()
		setIsCreated(t, *set, err)
	})

	t.Run("create set without lift_id", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		set.LiftId = 0
		err = set.Create()
		assertError(t, err, ErrSetMissingField)
	})

	t.Run("create set without created_at", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		set.CreatedAt = time.Time{}
		err = set.Create()
		assertError(t, err, ErrSetMissingField)
	})
}

func TestSetRetrieve(t *testing.T) {
	t.Parallel()
	t.Run("retrieve set", func(t *testing.T) {
		set, err := getTestSet()
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

func TestSetUpdate(t *testing.T) {
	set, err := getTestSet()
	assertNoError(t, err)

	err = set.Create()
	setIsCreated(t, *set, err)

	set.Done = true
	err = set.Update()
	assertNoError(t, err)

	retrievedSet, err := GetSet(set.Id)
	assertNoError(t, err)
	if retrievedSet.Done != true {
		t.Error("update failed: expected done to be true")
	}
}

func TestSetDelete(t *testing.T) {
	t.Parallel()
	set, err := getTestSet()
	assertNoError(t, err)
	set.Create()
	setIsCreated(t, *set, err)

	set.Delete()

	_, err = GetSet(set.Id)
	assertError(t, err, err)

}

func TestSetQuantityCreate(t *testing.T) {
	t.Parallel()
	t.Run("create SetQuantity", func(t *testing.T) {
		set, err := getTestSet()
		assertNoError(t, err)
		set.Create()
		setIsCreated(t, *set, err)

		sq := getTestSetQuantity(set)

		err = sq.Create()
		setQuantityIscreated(t, *sq, err)
	})

	t.Run("create SetQuantity without foreign key", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		set.Create()
		setIsCreated(t, *set, err)

		sq := getTestSetQuantity(set)
		sq.SetId = 0
		err = sq.Create()
		assertError(t, err, ErrSetQuantityMissingField)
	})

	t.Run("create SetQuantity with missing field (should fail)", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		set.Create()
		setIsCreated(t, *set, err)

		sq := getTestSetQuantity(set)
		sq.Reptype = 0
		err = sq.Create()
		assertError(t, err, ErrSetQuantityMissingField)
	})

	t.Run("create SetQuantity with missing field (should succeed)", func(t *testing.T) {
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		set.Create()
		setIsCreated(t, *set, err)

		sq := getTestSetQuantity(set)
		sq.Weight = 0
		err = sq.Create()
		setQuantityIscreated(t, *sq, err)
	})
}

func TestSetquantityRetrieve(t *testing.T) {
	t.Parallel()
	t.Run("retrieving SetQuantity", func(t *testing.T) {
		// setup set quantity
		t.Parallel()
		set, err := getTestSet()
		assertNoError(t, err)
		err = set.Create()
		setIsCreated(t, *set, err)
		sq := getTestSetQuantity(set)
		err = sq.Create()
		setQuantityIscreated(t, *sq, err)

		retrievedSq, err := GetSetQuantity(sq.Id)
		assertNoError(t, err)

		if retrievedSq.Id != sq.Id {
			t.Errorf("expected SetQuantity with id %d, got SetQuantity with id %d", sq.Id, retrievedSq.Id)
		}
	})
	t.Run("retrieving SetQuantity that doesn't exist", func(t *testing.T) {
		t.Parallel()
		_, err := GetSetQuantity(-1)
		assertError(t, err, err)
	})

}

// returns set struct with populated fields
// creates arbitrary lift for parent
func getTestSet() (set *Set, err error) {
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

func getTestSetQuantity(set *Set) (sq *SetQuantity) {
	sq = &SetQuantity{
		SetId:        set.Id,
		Reptype:      Count,
		Quantity:     8,
		Weight:       60,
		PlannedRatio: 70,
		Ratiotype:    Percentage,
		CreatedAt:    time.Now(),
	}

	return

}

func setIsCreated(t *testing.T, set Set, err error) {
	assertNoError(t, err)
	if set.Id == 0 {
		t.Error("Insersion failed: set id is still 0")
	}
}

func setQuantityIscreated(t *testing.T, sq SetQuantity, err error) {
	assertNoError(t, err)
	if sq.Id == 0 {
		t.Error("Insersion failed: SetQuantity id is still 0")
	}
}
