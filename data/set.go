package data

import "time"

type Set struct {
	Id        int
	LiftId    int `db:"lift_id"`
	WorkoutId int `db:"workout_id"`
	Done      bool
	CreatedAt time.Time `db:"created_at"`
}

// SetId, Reptype, Reptype and createdAt are necessary fields
type SetQuantity struct {
	Id           int
	SetId        int     `db:"set_id"`
	Reptype      RepType `db:"rep_type"`
	Quantity     int
	Weight       float64
	PlannedRatio int       `db:"planned_ratio"`
	Ratiotype    RatioType `db:"ratio_type"`
	CreatedAt    time.Time `db:"created_at"`
}

type RepType int64
type RatioType int64

const (
	Duration RepType = iota
	Count    RepType = iota
)

const (
	REM        RatioType = iota
	Percentage RatioType = iota
)

const (
	ErrSetMissingField         = SetErr("Set struct is missing field")
	ErrSetQuantityMissingField = SetErr("Set Quantity struct is missing field")
)

type SetErr string

func (err SetErr) Error() string {
	return string(err)
}

func (set *Set) Create() (err error) {
	if set.LiftId == 0 || set.WorkoutId == 0 || set.CreatedAt.IsZero() {
		return ErrSetMissingField
	}

	statement := "insert into sets (lift_id, workout_id, done, created_at) values ($1, $2, $3, $4) returning id"

	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	err = stmt.QueryRow(set.LiftId, set.WorkoutId, set.Done, set.CreatedAt).Scan(&set.Id)
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

func (set *Set) Delete() (err error) {
	_, err = db.Exec("delete from sets where id = $1", set.Id)
	return
}

func (sq *SetQuantity) Create() (err error) {
	if sq.SetId == 0 || sq.Reptype == 0 || sq.CreatedAt.IsZero() {
		return ErrSetQuantityMissingField
	}

	statement := "insert into setquantities (set_id, rep_type, quantity, weight, planned_ratio, ratio_type, created_at) values ($1, $2, $3, $4, $5, $6, $7) returning id"

	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	stmt.QueryRow(sq.SetId, sq.Reptype, sq.Quantity, sq.Weight, sq.PlannedRatio, sq.Ratiotype, sq.CreatedAt).Scan(&sq.Id)

	return
}

func GetSetQuantity(id int) (sq SetQuantity, err error) {
	err = db.QueryRowx("select * from setquantities where id = $1", id).StructScan(&sq)
	return
}

func (sq *SetQuantity) Update() (err error) {
	_, err = db.Exec("update setquantities set rep_type = $2, quantity = $3, weight = $4, planned_ratio = $5, ratio_type = $6 where id = $1", sq.Id, sq.Reptype, sq.Quantity, sq.Weight, sq.PlannedRatio, sq.Ratiotype)
	return
}

func (sq *SetQuantity) Delete() (err error) {
	_, err = db.Exec("delete from setquantities where id = $1", sq.Id)
	return
}

func GetSetInfos(workoutId, liftId int) (setinfos []SetInfo, err error) {

	rows, err := db.Queryx("select * from sets where workout_id=$1 and lift_id=$2", workoutId, liftId)

	defer func() {
		if rows.Err() != nil {
			err = rows.Err()
		}
		rows.Close()
	}()

	for rows.Next() {
		var set Set
		err = rows.StructScan(&set)
		if err != nil {
			rows.Close()
			return
		}

		var sq SetQuantity
		err = db.QueryRowx("select * from setquantities where set_id=$1", set.Id).StructScan(&sq)
		if err != nil {
			rows.Close()
			return
		}

		setinfos = append(setinfos, SetInfo{Set: set, Quantity: sq})
	}

	return
}
