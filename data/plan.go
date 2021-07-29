package data

import "time"

type Plan struct {
	Id        int
	Name      string
	Duration  int
	Frequency int
	CreatedAt time.Time `db:"created_at"`
}

const (
	ErrPlanMissingField    = PlanErr("Plan struct missing field")
	ErrPlanRetreiveFailure = PlanErr("Couldn't retreive Plan")
)

type PlanErr string

func (e PlanErr) Error() string {
	return string(e)
}

func (plan *Plan) Create() (err error) {
	// check for empty fields
	if plan.Name == "" || plan.Duration == 0 || plan.Frequency == 0 || plan.CreatedAt.IsZero() {
		return ErrPlanMissingField
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

func (plan *Plan) Delete() (err error) {
	_, err = db.Exec("delete from plans where id = $1", plan.Id)
	return
}
