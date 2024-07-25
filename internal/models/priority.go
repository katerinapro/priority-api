package models

import (
	"time"
)

type Priority struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created_date"`
}
