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

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables")
	}
	if local {
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
	if err != nil {
		log.Fatalf("error opening %s, %s", file, err)
	}
	schema = string(schemaStream)

	return
}

const (
	ErrMissingField    = PlanErr("Plan struct missing field")
	ErrRetreiveFailure = PlanErr("Couldn't retreive Plan")
)

type PlanErr string

func (e PlanErr) Error() string {
	return string(e)
}

func (plan *Plan) Create() (err error) {
	// check for empty fields
	if plan.Name == "" || plan.Duration == 0 || plan.Frequency == 0 || plan.CreatedAt.IsZero() {
		return ErrMissingField
	}

	statement := "insert into plans (name, duration, frequency, created_at) values ($1, $2, $3, $4) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(plan.Name, plan.Duration, plan.Frequency, plan.CreatedAt).Scan(&plan.Id)
	return
}

func GetPlan(id int) (plan Plan, err error) {
	err = db.QueryRowx("select * from plans where id=$1", id).StructScan(&plan)
	return
}

func (plan *Plan) Update() (err error) {
	_, err = db.Exec("update plans set name = $2, duration = $3, frequency = $4 where id = $1", plan.Id, plan.Name, plan.Duration, plan.Frequency)
	return
}
