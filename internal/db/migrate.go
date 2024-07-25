package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

type Migration struct {
	Name string
	Path string
}

func RunMigrations(db *sql.DB, direction string) error {
	// Ensure schema_migrations table exists
	if err := ensureSchemaMigrationsTable(db); err != nil {
		return err
	}

	migrations, err := getMigrations("migrations", direction)
	if err != nil {
		log.Fatal(err)
	}

	for _, migration := range migrations {
		if direction == "up" {
			if migrationApplied(db, migration.Name) {
				fmt.Printf("Skipping already applied migration: %s\n", migration.Name)
				continue
			}
		}

		fmt.Printf("Processing migration: %s\n", migration.Name)
		if err := executeMigration(db, migration.Path); err != nil {
			log.Fatalf("Failed to execute migration %s: %v", migration.Name, err)
		}

		if direction == "up" {
			if err := recordMigration(db, migration.Name); err != nil {
				log.Fatalf("Failed to record migration %s: %v", migration.Name, err)
			}
		} else if direction == "down" {
			if err := deleteMigrationRecord(db, migration.Name); err != nil {
				log.Fatalf("Failed to delete migration record %s: %v", migration.Name, err)
			}
		}
	}

	if direction == "down" {
		if err := deleteSchemaMigrationsTable(db); err != nil {
			return err
		}
	}

	return nil
}

func ensureSchemaMigrationsTable(db *sql.DB) error {
	query := `
	CREATE SCHEMA IF NOT EXISTS lo;
	CREATE TABLE IF NOT EXISTS lo.schema_migrations (
		version VARCHAR(255) PRIMARY KEY
	);`
	_, err := db.Exec(query)
	return err
}

func deleteSchemaMigrationsTable(db *sql.DB) error {
	query := `
	DROP TABLE lo.schema_migrations;`
	_, err := db.Exec(query)
	return err
}

func getMigrations(dir string, direction string) ([]Migration, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var migrations []Migration
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), fmt.Sprintf(".%s.", direction)) {
			migrations = append(migrations, Migration{
				Name: file.Name(),
				Path: dir + "/" + file.Name(),
			})
		}
	}

	sortMigrations(migrations, direction)
	return migrations, nil
}

func sortMigrations(migrations []Migration, direction string) {
	sort.Slice(migrations, func(i, j int) bool {
		if direction == "up" {
			return migrations[i].Name < migrations[j].Name
		}
		return migrations[i].Name > migrations[j].Name
	})
}

func executeMigration(db *sql.DB, filePath string) error {
	query, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(query))
	return err
}

func migrationApplied(db *sql.DB, migrationName string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)", migrationName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if migration is applied: %v", err)
	}
	return exists
}

func recordMigration(db *sql.DB, migrationName string) error {
	_, err := db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", migrationName)
	return err
}

func deleteMigrationRecord(db *sql.DB, migrationName string) error {
	_, err := db.Exec("DELETE FROM schema_migrations WHERE version = $1", migrationName)
	return err
}
