package data

import "time"

type Workout struct {
	Id        int
	PlanId    int
	Name      string
	WeekNo    int
	Date      time.Time
	CreatedAt time.Time
}
