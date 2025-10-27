// backend/migrations/migrate.go

package migrations

import (
	"log"
	"database/sql"
	"github.com/kobekimmes/nyab/backend/db"
)


func RunMigrationsUp() {
	db.DbInit()
	defer db.Db.Close()

	applied := getAppliedMigrations(db.Db)

	for _, migration := range All {
		if isMigrationApplied(migration.Name, applied) {
			continue
		}

		log.Printf("Applying migration: %s\n", migration.Name)
		if err := migration.Up(db.Db); err != nil {
			log.Fatalf("Failed to apply migration %s: %v", migration.Name, err)
		}

		markMigrationApplied(migration.Name, db.Db)
		log.Printf("Migration applied: %s\n", migration.Name)
	}
}

func RunMigrationsDown() {
	db.DbInit()
	defer db.Db.Close()

	applied := getAppliedMigrations(db.Db)

	for _, migration := range All {
		if !isMigrationApplied(migration.Name, applied) {
			continue
		}

		log.Printf("Rolling back migration: %s\n", migration.Name)
		if err := migration.Down(db.Db); err != nil {
			log.Fatalf("Failed to rollback migration %s: %v", migration.Name, err)
		}

		unmarkMigration(migration.Name, db.Db)
		log.Printf("Migration rolled back: %s\n", migration.Name)
	}
}


func getAppliedMigrations(db *sql.DB) map[string]bool {

	rows, err := db.Query("SELECT name FROM migrations")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	applied := map[string]bool{}
	for rows.Next() {
		var name string
		rows.Scan(&name)

		applied[name] = true
	}

	return applied

}

func isMigrationApplied(migration_name string, applied_migrations map[string]bool) bool {
	return applied_migrations[migration_name]
}

func markMigrationApplied(migration_name string, db *sql.DB) {
	_, err := db.Exec("INSERT INTO migrations (name) VALUES ($1)", migration_name)

	if err != nil {
		log.Fatal(err)
	}
}

func unmarkMigration(migration_name string, db *sql.DB) {
	_, err := db.Exec("DELETE FROM migrations WHERE name=$1", migration_name)
	if err != nil {
		log.Fatalf("Failed to unmark migration %s: %v", migration_name, err)
	}
}



