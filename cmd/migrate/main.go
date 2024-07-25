package main

import (
	"database/sql"
	"log"
	"os"

	//"os"

	"github.com/katerinapro/priority-api/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	database.Exec(`set search_path='lo'`)
	defer database.Close()

	if err := db.RunMigrations(database, "up"); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}
