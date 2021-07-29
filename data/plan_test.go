package data

import (
	"testing"
	"time"
)

func TestPlanCreate(t *testing.T) {
	t.Parallel()
	t.Run("creating plan", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Duration:  4,
			Frequency: 3,
			CreatedAt: time.Now(),
		}

		err := plan.Create()
		assertNoError(t, err)
		if plan.Id == 0 {
			t.Errorf("insertion failed: plan id is still %d", plan.Id)
		}
	})

	t.Run("creating plan with no name", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Duration:  4,
			Frequency: 3,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testPlanEmptyField(t, plan, err)
	})

	t.Run("creating plan with no duration", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Frequency: 3,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testPlanEmptyField(t, plan, err)
	})

	t.Run("creating plan with no frequency", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Duration:  4,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testPlanEmptyField(t, plan, err)
	})

	t.Run("creating plan with no CreatedAt", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Duration:  4,
			Frequency: 3,
		}
		err := plan.Create()
		testPlanEmptyField(t, plan, err)
	})
}

func testPlanEmptyField(t *testing.T, plan Plan, err error) {
	t.Helper()
	assertError(t, err, ErrPlanMissingField)
	if plan.Id != 0 {
		t.Errorf("error: insertion did not fail")
	}
}

func TestPlanRetrieve(t *testing.T) {
	t.Parallel()
	createdAt, err := time.Parse(time.RFC822, "21 Jul 28 14:11 CDT")
	assertNoError(t, err)

	plan := Plan{
		Name:      "Retrieve test plan name",
		Duration:  100,
		Frequency: 80,
		CreatedAt: createdAt,
	}
	plan.Create()

	retreivedPlan, err := GetPlan(plan.Id)
	assertNoError(t, err)
	if retreivedPlan.Id != plan.Id {
		t.Errorf("wrong workout retrieved: retrieved plan with id %d, wanted plan with id %d", retreivedPlan.Id, plan.Id)
	}
}

func TestPlanUpdate(t *testing.T) {
	t.Parallel()
	plan := Plan{
		Name:      "pre-update test plan name",
		Duration:  111,
		Frequency: 88,
		CreatedAt: time.Now(),
	}
	plan.Create()

	preUpdatedPlan, err := GetPlan(plan.Id)
	assertNoError(t, err)

	preUpdatedPlan.Name = "Updated test plan name"
	err = preUpdatedPlan.Update()
	assertNoError(t, err)

	updatedPlan, err := GetPlan(plan.Id)
	assertNoError(t, err)

	if updatedPlan.Name != "Updated test plan name" {
		t.Errorf("field not updated")
	}
}

func TestPlanDelete(t *testing.T) {
	t.Parallel()
	plan := Plan{
		Name:      "pre-update test plan name",
		Duration:  111,
		Frequency: 88,
		CreatedAt: time.Now(),
	}
	plan.Create()

	toDeletePlan, err := GetPlan(plan.Id)
	assertNoError(t, err)

	toDeletePlan.Delete()

	deletedPlan, err := GetPlan(toDeletePlan.Id)
	assertError(t, err, err)

	if deletedPlan.Id != 0 {
		t.Error("Plan has not been deleted after delete function")
	}
}
