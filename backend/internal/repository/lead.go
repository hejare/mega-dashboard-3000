package repository

import "database/sql"

type Lead struct {
	Id           int
	Organization string
	Stack        []string
	Role         string
	Contact      string
	ConsultantId int
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
