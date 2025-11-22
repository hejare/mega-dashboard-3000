package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type DashboardRepository struct {
	db *sql.DB
}

type DashboardRow struct {
	ConsultantId           int
	ConsultantName         string
	AssignmentId           *int
	AssignmentEndDate      *time.Time
	AssignmentOrganization *string
	ExtensionStatus        string
	ChangeStatus           string
	Leads                  []Lead
}

func NewDashboardRepository(db *sql.DB) *DashboardRepository {
	return &DashboardRepository{
		db: db,
	}
}

func (r *DashboardRepository) GetDashboard() ([]DashboardRow, error) {
	query := `
	SELECT
    c.id AS consultant_id,
    c.name AS consultant_name,
    c.change_status AS consultant_change_status,
    c.probable_extension_status AS consultant_probable_extension_status,
    a.id  AS assignment_id,
    a.period_end_at AS assignment_period_end_at,
    a.organization AS assignment_organization,
    COALESCE(
        json_agg(
            jsonb_build_object(
                'Id', l.id,
                'Organization', l.organization,
                'Stack', array_to_json(l.stack)
            )
        ) FILTER (WHERE l.id IS NOT NULL),
        '[]'
    ) AS leads
FROM consultants c
LEFT JOIN assignments a ON a.consultant_id = c.id AND NOW() BETWEEN a.period_start_at AND a.period_end_at
LEFT JOIN leads l ON l.consultant_id = c.id
GROUP BY c.id, c.name, c.change_status, c.probable_extension_status, a.id
ORDER BY c.id;
`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dashboard []DashboardRow

	for rows.Next() {
		var row DashboardRow
		var leadsJSON []byte

		err := rows.Scan(
			&row.ConsultantId,
			&row.ConsultantName,
			&row.ChangeStatus,
			&row.ExtensionStatus,
			&row.AssignmentId,
			&row.AssignmentEndDate,
			&row.AssignmentOrganization,
			&leadsJSON,
		)
		if err != nil {
			fmt.Printf("\n\nERRROR: %v\n\n", err)
			return nil, err
		}

		if err := json.Unmarshal(leadsJSON, &row.Leads); err != nil {
			return nil, err
		}
		dashboard = append(dashboard, row)
	}

	return dashboard, nil

}
