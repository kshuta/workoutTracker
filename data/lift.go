package data

import "time"

type Lift struct {
	Id        int
	Name      string
	Max       float64
	IsDeleted bool      `db:"is_deleted"`
	CreatedAt time.Time `db:"created_at"`
}

const (
	ErrLiftMissingField = LiftErr("Lift struct is missing a field")
)

type LiftErr string

func (l LiftErr) Error() string {
	return string(l)
}

func (lift *Lift) Create() (err error) {
	lift.CreatedAt = time.Now()
	lift.IsDeleted = false

	statement := "insert into lifts (name, max, created_at, is_deleted) values ($1, $2, $3, $4) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(lift.Name, lift.Max, lift.CreatedAt, lift.IsDeleted).Scan(&lift.Id)
	return
}

func GetLift(id int) (lift Lift, err error) {
	err = db.QueryRowx("select * from lifts where id = $1 and is_deleted=false", id).StructScan(&lift)
	return
}

func GetLifts() (lifts []Lift, err error) {
	rows, err := db.Queryx("select * from lifts where is_deleted=false")
	if err != nil {
		rows.Close()
		return
	}

	for rows.Next() {
		var lift Lift
		rows.StructScan(&lift)
		lifts = append(lifts, lift)
	}

	rows.Close()
	err = rows.Err()
	return
}

func (lift *Lift) Update() (err error) {
	_, err = db.Exec("update lifts set name = $2, max = $3 where id = $1", lift.Id, lift.Name, lift.Max)
	return
}

func (lift *Lift) Delete() (err error) {
	_, err = db.Exec("update lifts set is_deleted = true where id = $1", lift.Id)
	return
}

// function to delete all lifts, only used within tests
func deleteAllLifts() (err error) {
	_, err = db.Exec("update lifts set is_deleted = true")
	return
}
