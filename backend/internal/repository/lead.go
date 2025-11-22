package repository

import (
	"database/sql"
	"fmt"
)

type Lead struct {
	Id           int      `json:"id"`
	Stack        []string `json:"stack"`
	Role         string   `json:"role"`
	Contact      string   `json:"contact"`
	ConsultantId int      `json:"consultantId"`
	Organization string   `json:"organization"`
	Title        *string  `json:"title"`
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

type LeadWithConsultantID struct {
	Lead
	ConsultantName string `json:"consultant_name"`
}

func (r *LeadRepository) Search(query string) ([]LeadWithConsultantID, error) {
	processedQuery := query + "%"

	sqlQuery := `
		SELECT leads.id, leads.organization, consultants.name
		FROM leads
		LEFT JOIN consultant_id_to_lead_id_mapping
			ON leads.id = consultant_id_to_lead_id_mapping.lead_id
		LEFT JOIN consultants
			ON consultant_id_to_lead_id_mapping.consultant_id = consultants.id
		WHERE organization ILIKE $1; 
	`

	rows, err := r.db.Query(sqlQuery, processedQuery)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var leads []LeadWithConsultantID

	for rows.Next() {
		var l LeadWithConsultantID
		var stack []string
		if err := rows.Scan(&l.Id, &l.Organization, &l.ConsultantName); err != nil {
			fmt.Print("Error scanning row: ", err)
			return nil, err
		}
		l.Stack = stack
		leads = append(leads, l)
	}
	return leads, nil
}
