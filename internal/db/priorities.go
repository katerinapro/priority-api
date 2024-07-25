package db

import (
	"github.com/katerinapro/priority-api/internal/models"
	_ "github.com/lib/pq"
)

func GetPriorities() ([]models.Priority, error) {
	rows, err := db.Query("SELECT * FROM lo.get_all_priorities()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var priorities []models.Priority
	for rows.Next() {
		var p models.Priority
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Created); err != nil {
			return nil, err
		}
		priorities = append(priorities, p)
	}

	return priorities, nil
}
