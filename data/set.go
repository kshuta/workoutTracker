package data

import "time"

type Set struct {
	Id        int
	LiftId    int `db:"lift_id"`
	Done      bool
	CreatedAt time.Time `db:"created_at"`
}

type SetQuantity struct {
	Id           int
	SetId        int    `db:"set_id"`
	RepType      string `db:"rep_type"`
	Quantity     int
	Weight       float64
	PlannedRatio int       `db:"planned_ratio"`
	Ratiotype    string    `db:"ratio_type"`
	CreatedAt    time.Time `db:"created_at"`
}

const (
	ErrSetMissingField = SetErr("Set struct is missing field")
)

type SetErr string

func (err SetErr) Error() string {
	return string(err)
}

func (set *Set) Create() (err error) {
	if set.LiftId == 0 || set.CreatedAt.IsZero() {
		return ErrSetMissingField
	}

	statement := "insert into sets (lift_id, done, created_at) values ($1, $2, $3) returning id"

	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	err = stmt.QueryRow(set.LiftId, set.Done, set.CreatedAt).Scan(&set.Id)
	return
}

func GetSet(id int) (set Set, err error) {
	err = db.QueryRowx("select * from sets where id = $1", id).StructScan(&set)
	return
}

func (set *Set) Update() (err error) {
	_, err = db.Exec("update sets set done = $1", set.Done)
	return
}
