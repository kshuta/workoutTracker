package data

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	if local {
		err = godotenv.Load(".env.local")
		if err != nil {
			log.Fatalln("Failed to load environment variables")
		}
		DSN := fmt.Sprintf(dsnUrlFormat, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("POSTGRES_DB"))
		db, err = sqlx.Connect("postgres", DSN)
		check(err)
		if err != nil {
			log.Fatalln("couldn't connect to database: ", err)
		}
	}

	schema, err := getSQL(initSchemaFile)

	check(err)
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln("couldn't execute schema: ", err)
	}

	log.Println("Test set up complete")

}

// Tests if the given error matches the expected error.
// if identical errors are passed in, the function can be used
// to test if an error exists
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("error not envoked")
	}

	if got != want {
		t.Fatalf("expected error %q, got %q", want, got)
	}
}

// checks that the given error is nil
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("error found: ", err)
	}
}
