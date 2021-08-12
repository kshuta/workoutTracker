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

func GetLift(id int) (lift Lift, err error) {
	err = db.QueryRowx("select * from lifts where id = $1", id).StructScan(&lift)
	return
}

func (lift *Lift) Update() (err error) {
	_, err = db.Exec("update lifts set name = $2, max = $3 where id = $1", lift.Id, lift.Name, lift.Max)
	return
}

func (lift *Lift) Delete() (err error) {
	_, err = db.Exec("delete from lifts where id = $1", lift.Id)
	return
}
