package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const dsnUrlFormat = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
const local = true

// var remoteDbDSN =

var initSchemaFile = "setup.sql"
var db *sqlx.DB

// checks if error is nil
func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	if local {
		err = godotenv.Load(".env.local")
		if err != nil {
			log.Fatalln("Failed to load environment variables: ", err)
		}
		DSN := fmt.Sprintf(dsnUrlFormat, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("POSTGRES_DB"))
		db, err = sqlx.Connect("postgres", DSN)
		check(err)
	}

	// get sql statement with schema
	// and execute it.
	schema, err := getSQL(initSchemaFile)
	check(err)
	db.MustExec(schema)
}

// Reads the file with the passed in filename, and
// returns sql statement within the file.
// file extension must be .sql
func getSQL(file string) (schema string, err error) {
	if filepath.Ext(file) != ".sql" {
		err = errors.New("incorrect file extension")
		return
	}
	schemaStream, err := ioutil.ReadFile(file)
	// temp err
	if err != nil && file == "setup.sql" {
		schemaStream, err = ioutil.ReadFile("data/setup.sql")
	}
	if err != nil {
		log.Fatalf("error opening %s, %s", file, err)
	}
	schema = string(schemaStream)

	return
}
