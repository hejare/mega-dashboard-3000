package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Lead struct {
	Id           int
	Organization *string
	Stack        []string
	Role         *string
	Contact      *string
	ConsultantId int
	Title        *string
}

type LeadRepository struct {
	db *sql.DB
}

type CreateLeadData struct {
	Organization  string
	Contact       string
	Role          string
	Stack         []string
	Title         string
	ConsultantIds []int
}

func NewLeadRepository(db *sql.DB) *LeadRepository {
	return &LeadRepository{
		db: db,
	}
}

func (r *LeadRepository) CreateLead(data *CreateLeadData) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var leadID int
	err = tx.QueryRow(
		`INSERT INTO leads (organization, contact, role, stack, title)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		data.Organization, data.Contact, data.Role, data.Stack, data.Title,
	).Scan(&leadID)
	if err != nil {
		return err
	}

	for _, consultantID := range data.ConsultantIds {
		_, err := tx.Exec(
			`INSERT INTO consultant_id_to_lead_id_mapping (consultant_id, lead_id)
			VALUES ($1, $2)`,
			consultantID, leadID,
		)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *LeadRepository) Search(query string) ([]Lead, error) {
	sqlQuery := `
	SELECT id, organization, role, stack, title
			FROM leads
			WHERE to_tsvector('simple', coalesce(role,'') || ' ' || array_to_string(stack,' '))
			      @@ to_tsquery('simple', $1)
			ORDER BY ts_rank(
			    to_tsvector('simple', coalesce(role,'') || ' ' || array_to_string(stack,' ')),
			    to_tsquery('simple', $1)
			) DESC
			LIMIT 50
	`

	fmt.Printf("query string: %s\n", query)

	rows, err := r.db.Query(sqlQuery, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leads []Lead
	for rows.Next() {
		var l Lead
		var stack []string
		if err := rows.Scan(&l.Id, &l.Organization, &l.Role, pq.Array(&stack), &l.Title); err != nil {
			fmt.Printf("failed scanning: %v\n", err)
			return nil, err
		}
		l.Stack = stack
		leads = append(leads, l)
	}
	return leads, nil
}
