package data

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Tests if the given error matches the expecte error.
// if identical errors are passed in, the function can be used
// to test if an error exists
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("error not envoked")
	}
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("error found: ", err)
	}
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables")
	}
	if local {
		DSN := fmt.Sprintf(dsnUrlFormat, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("POSTGRES_DB"))
		db, err = sqlx.Connect("postgres", DSN)
		check(err)
		if err != nil {
			log.Fatalln("couldn't connect to database: ", err)
		}
	}

	schema, err := getSQL(initSchemaFile)
	if err != nil {
		log.Fatalln("getSQL failed: ", err)
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln("couldn't execute schema: ", err)
	}

	log.Println("Test set up complete")

}
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
		testEmptyField(t, plan, err)
	})

	t.Run("creating plan with no duration", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Frequency: 3,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testEmptyField(t, plan, err)
	})

	t.Run("creating plan with no frequency", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Duration:  4,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testEmptyField(t, plan, err)
	})

	t.Run("creating plan with no CreatedAt", func(t *testing.T) {
		t.Parallel()
		plan := Plan{
			Name:      "Test Plan Name",
			Duration:  4,
			Frequency: 3,
		}
		err := plan.Create()
		testEmptyField(t, plan, err)
	})
}

func testEmptyField(t *testing.T, plan Plan, err error) {
	t.Helper()
	assertError(t, err, ErrMissingField)
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
		t.Errorf("Retrieved plan with id %d, wanted plan with id %d", retreivedPlan.Id, plan.Id)
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
