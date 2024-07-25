package db

import (
	"database/sql"
	"fmt"

	"github.com/katerinapro/priority-api/internal/models"
	_ "github.com/lib/pq"
)

func GetPriority(id string) (models.Priority, error) {
	var priority models.Priority
	err := db.QueryRow("SELECT * FROM lo.get_priority_by_id($1)", id).Scan(
		&priority.ID,
		&priority.Title,
		&priority.Description,
		&priority.Created,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Priority{}, fmt.Errorf("no priority found with id %s", id)
		}
		return models.Priority{}, fmt.Errorf("error querying priority: %v", err)
	}

	return priority, nil
}

func CreatePriority(priority *models.Priority) error {
	err := db.QueryRow(
		"INSERT INTO lo.priorities (title, description) VALUES ($1, $2) RETURNING id, created_date",
		priority.Title, priority.Description,
	).Scan(&priority.ID, &priority.Created)

	if err != nil {
		return fmt.Errorf("error creating priority: %v", err)
	}

	return nil
}

func UpdatePriority(id string, priority *models.Priority) error {
	_, err := db.Exec(
		"UPDATE lo.priorities SET title=$1, description=$2 WHERE id=$3",
		priority.Title, priority.Description, id,
	)

	if err != nil {
		return fmt.Errorf("error updating priority: %v", err)
	}

	err = db.QueryRow(
		"SELECT id, title, description, created_date FROM lo.priorities WHERE id=$1",
		id,
	).Scan(&priority.ID, &priority.Title, &priority.Description, &priority.Created)

	if err != nil {
		return fmt.Errorf("error retrieving updated priority: %v", err)
	}

	return nil
}

func DeletePriority(id string) error {
	_, err := db.Exec("DELETE FROM lo.priorities WHERE id=$1", id)

	if err != nil {
		return fmt.Errorf("error deleting priority: %v", err)
	}

	return nil
}
