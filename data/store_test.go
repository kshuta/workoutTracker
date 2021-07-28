package data

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	var err error
	db, err = sqlx.Connect("postgres", loaclDbDSN)
	if err != nil {
		log.Fatalln("couldn't connect to database: ", err)
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
	t.Run("creating plan", func(t *testing.T) {
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
		plan := Plan{
			Duration:  4,
			Frequency: 3,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testEmptyField(t, plan, err)
	})

	t.Run("creating plan with no duration", func(t *testing.T) {
		plan := Plan{
			Name:      "Test Plan Name",
			Frequency: 3,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testEmptyField(t, plan, err)
	})

	t.Run("creating plan with no frequency", func(t *testing.T) {
		plan := Plan{
			Name:      "Test Plan Name",
			Duration:  4,
			CreatedAt: time.Now(),
		}
		err := plan.Create()
		testEmptyField(t, plan, err)
	})

	t.Run("creating plan with no CreatedAt", func(t *testing.T) {
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
