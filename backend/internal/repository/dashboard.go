package repository

import "database/sql"

type DashboardRepository struct {
	db *sql.DB
}

type Dashboard struct {
}

func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{
		db: db,
	}
}

func (r *DashboardRepository) GetDashboard() {
}
