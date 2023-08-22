package pgsql

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var migrationsFs embed.FS

var migrations = &migrate.EmbedFileSystemMigrationSource{
	FileSystem: migrationsFs,
	Root:       "migrations",
}

func MigrateUp(db *sql.DB) error {
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %s", err.Error())
	}
	log.Printf("applied %d migrations", n)
	return nil
}

func MigrateDown(db *sql.DB) error {
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %s", err.Error())
	}
	log.Printf("applied %d migrations", n)
	return nil
}
