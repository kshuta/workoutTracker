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
