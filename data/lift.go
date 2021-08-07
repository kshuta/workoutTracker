package data

import "time"

type Lift struct {
	Id        int
	Name      string
	Max       float64
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
	if lift.CreatedAt.IsZero() {
		err = ErrLiftMissingField
		return
	}

	statement := "insert into lifts (name, max, created_at) values ($1, $2, $3) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(lift.Name, lift.Max, lift.CreatedAt).Scan(&lift.Id)
	return
}
