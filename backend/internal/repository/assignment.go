package repository

import "time"

type Assignment struct {
	Id            int
	Company       string
	Stack         []string
	Role          string
	Contact       string
	HourlyPrice   int
	PeriodStartAt time.Time
	PeriodEndAt   time.Time
	Lead          *Lead
	Consultant    *Consultant
}
