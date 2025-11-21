package repository

import "database/sql"

type ConsultantRepository struct {
	db *sql.DB
}

type Consultant struct {
	Id           int
	Name         string
	ChangeStatus string
}

func NewConsultantRepository(db *sql.DB) *ConsultantRepository {
	return &ConsultantRepository{
		db: db,
	}
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
