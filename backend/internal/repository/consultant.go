package repository

import "database/sql"

type ConsultantRepository struct {
	db *sql.DB
}

type Consultant struct {
	Id                      int
	Name                    string
	ChangeStatus            string
	ProbableExtensionStatus string
}

func NewConsultantRepository(db *sql.DB) *ConsultantRepository {
	return &ConsultantRepository{
		db: db,
	}
}

func (r *ConsultantRepository) GetConsultant(id int) (*Consultant, error) {
	row := r.db.QueryRow("SELECT id, name, change_status FROM consultants WHERE id=$1", id)
	var c Consultant
	if err := row.Scan(&c.Id, &c.Name, &c.ChangeStatus); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *ConsultantRepository) GetAllConsultants() ([]Consultant, error) {
	rows, err := r.db.Query("SELECT id, name, change_status FROM consultants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consultants []Consultant
	for rows.Next() {
		var c Consultant
		if err := rows.Scan(&c.Id, &c.Name, &c.ChangeStatus); err != nil {
			return nil, err
		}
		consultants = append(consultants, c)
	}
	return consultants, nil
}
