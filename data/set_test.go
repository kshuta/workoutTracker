package data

import (
	"database/sql"
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
		assertError(t, err, sql.ErrNoRows)
	})

}

func TestSetQuantityUpdate(t *testing.T) {
	t.Parallel()
	set, err := getTestSet()
	assertNoError(t, err)
	err = set.Create()
	setIsCreated(t, *set, err)
	sq := getTestSetQuantity(set)
	err = sq.Create()
	setQuantityIscreated(t, *sq, err)

	updatedWeight := 55.5
	sq.Weight = updatedWeight
	sq.Update()

	retrievedSq, err := GetSetQuantity(sq.Id)
	assertNoError(t, err)

	if retrievedSq.Weight != updatedWeight {
		t.Error("update error: field not updated")
	}
}

func TestSetQuantityDelete(t *testing.T) {
	t.Parallel()
	set, err := getTestSet()
	assertNoError(t, err)
	err = set.Create()
	setIsCreated(t, *set, err)
	sq := getTestSetQuantity(set)
	err = sq.Create()
	setQuantityIscreated(t, *sq, err)

	sq.Delete()

	_, err = GetSetQuantity(sq.Id)
	assertError(t, err, sql.ErrNoRows)

}

func TestGetSetInfos(t *testing.T) {
	t.Parallel()
	workout, lift, err := getSetParents()
	assertNoError(t, err)

	setinfos, err := creatTestSetInfos(workout, lift)
	assertNoError(t, err)

	retrievedSetInfos, err := GetSetInfos(workout.Id, lift.Id)
	assertNoError(t, err)

	if len(retrievedSetInfos) != len(setinfos) {
		t.Fatalf("length of SetInfo is not the same. Expected %d setinfos, got %d setinfos", len(setinfos), len(retrievedSetInfos))
	}

	for idx, setinfo := range setinfos {
		if setinfo.Set.Id != retrievedSetInfos[idx].Set.Id {
			t.Fatalf("Retrieved setinfo differs. Expected setinfo with set id %d, got set id %d", setinfo.Set.Id, retrievedSetInfos[idx].Set.Id)
		}
		if setinfo.Quantity.Id != retrievedSetInfos[idx].Quantity.Id {
			t.Fatalf("Retrieved setinfo differs. Expected setinfo with SetQuantity id %d, got SetQuantity id %d", setinfo.Quantity.Id, retrievedSetInfos[idx].Quantity.Id)
		}
	}
}

// returns multiple set struct with same workout id
// creates arbitrary lift and workout for parent
func creatTestSetInfos(workout *Workout, lift *Lift) (setinfos []SetInfo, err error) {
	setinfos = make([]SetInfo, 0)
	for i := 0; i < 4; i++ {
		set := Set{
			LiftId:    lift.Id,
			WorkoutId: workout.Id,
			Done:      false,
			CreatedAt: time.Now(),
		}
		err = set.Create()
		if err != nil {
			return
		}

		sq := SetQuantity{
			SetId:        set.Id,
			Reptype:      Count,
			Quantity:     i + 4,
			PlannedRatio: int(70 + (i+1)*4),
			Ratiotype:    Percentage,
			CreatedAt:    time.Now(),
		}
		err = sq.Create()
		if err != nil {
			return
		}

		setinfo := SetInfo{
			Set:      set,
			Quantity: sq,
		}

		setinfos = append(setinfos, setinfo)
	}

	return

}

// returns set struct with populated fields
// creates arbitrary lift and workout for parent
func getTestSet() (set *Set, err error) {
	workout, lift, err := getSetParents()
	if err != nil {
		return
	}

	set = &Set{
		LiftId:    lift.Id,
		WorkoutId: workout.Id,
		Done:      false,
		CreatedAt: time.Now(),
	}

	return
}

// gets created lift and workout to create sets
func getSetParents() (workout *Workout, lift *Lift, err error) {
	lift = getTestLift("lift for set test")
	err = lift.Create()
	if err != nil {
		return
	}

	workout = getTestWorkout("workout for set test")
	err = workout.Create()
	if err != nil {
		return
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
	t.Helper()
	assertNoError(t, err)
	if set.Id == 0 {
		t.Error("Insersion failed: set id is still 0")
	}
}

func setQuantityIscreated(t *testing.T, sq SetQuantity, err error) {
	t.Helper()
	assertNoError(t, err)
	if sq.Id == 0 {
		t.Error("Insersion failed: SetQuantity id is still 0")
	}
}
