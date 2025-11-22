package repository

import (
	"database/sql"
	"time"
)

type Assignment struct {
	Id            int         `json:"id"`
	Organization  string      `json:"organization"`
	Stack         []string    `json:"stack"`
	Role          string      `json:"role"`
	Contact       string      `json:"contact"`
	HourlyPrice   int         `json:"hourlyPrice"`
	PeriodStartAt time.Time   `json:"periodStartAt"`
	PeriodEndAt   time.Time   `json:"periodEndAt"`
	Lead          *Lead       `json:"lead"`
	Consultant    *Consultant `json:"consultant"`
}

type CreateAssignmentData struct {
	Title         string
	Organization  string
	Contact       string
	Role          string
	Stack         []string
	HourlyPrice   int
	PeriodStartAt time.Time
	PeriodEndAt   time.Time
	LeadID        int
	ConsultantID  int
}

type AssignmentRepository struct {
	db *sql.DB
}

func NewAssignmentRepository(db *sql.DB) *AssignmentRepository {
	return &AssignmentRepository{
		db: db,
	}
}

func (r *AssignmentRepository) CreateAssignment(data *CreateAssignmentData) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var assignmentID int
	err = tx.QueryRow(
		`INSERT INTO leads (organization, contact, role, stack, price, period_start_at, period_end_at, lead_id, consultant_id, title)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`,
		data.Organization, data.Contact, data.Role, data.Stack, data.HourlyPrice, data.PeriodStartAt, data.PeriodEndAt, data.LeadID, data.ConsultantID, data.Title,
	).Scan(&assignmentID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
