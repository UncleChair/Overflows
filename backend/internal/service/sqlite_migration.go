package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	_ "github.com/mattn/go-sqlite3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func createMigrationHistoryTable(ctx context.Context, db gdb.DB) error {
	if ctx == nil {
		ctx = context.Background()
	}

	span := trace.SpanFromContext(ctx)
	if span == nil {
		ctx, span = otel.Tracer("sqlite-migration").Start(ctx, "createMigrationHistoryTable")
	}
	defer span.End()

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS migration_history (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        migration_name TEXT UNIQUE,
        executed_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	if _, err := db.Exec(ctx, createTableSQL); err != nil {
		return fmt.Errorf("failed to create migration history table: %w", err)
	}
	return nil
}
func migrationExecuted(ctx context.Context, db gdb.DB, migrationName string) (bool, error) {
	query := `SELECT COUNT(*) FROM migration_history WHERE migration_name = ?`
	count, err := db.GetValue(ctx, query, migrationName)
	if err != nil {
		return false, fmt.Errorf("failed to check migration status: %w", err)
	}
	return count.Int() > 0, nil
}

func recordMigration(ctx context.Context, db gdb.DB, migrationName string) error {
	insertSQL := `INSERT INTO migration_history (migration_name) VALUES (?)`
	_, err := db.Exec(ctx, insertSQL, migrationName)
	return err
}

func executeSQLFile(ctx context.Context, db gdb.DB, sqlFile string) error {
	var content []byte
	if resource := gres.Get(sqlFile); resource != nil {
		content = resource.Content()
	} else {
		absolutePath := gfile.RealPath(sqlFile)
		if absolutePath != "" {
			content = gfile.GetBytes(absolutePath)
		}
	}
	_, err := db.Exec(ctx, string(content))
	if err != nil {
		return fmt.Errorf("failed to execute SQL from file %s: %v", sqlFile, err)
	}

	fmt.Printf("Successfully executed %s\n", sqlFile)
	return nil
}

func executeMigrations(ctx context.Context, db gdb.DB, migrationFolder string) {
	var (
		resources []*gres.File
		files     []string
		err       error
	)
	resources = gres.ScanDirFile(migrationFolder, "*", false)
	if len(resources) > 0 {
		for _, file := range resources {
			files = append(files, file.Name())
		}
	} else {
		absolutePath := gfile.RealPath(migrationFolder)
		files, err = gfile.ScanDir(absolutePath, "*", false)
		if err != nil {
			log.Printf("Error scanning migration folder %s: %v\n", migrationFolder, err)
			return
		}
	}

	for _, filename := range files {
		if filepath.Ext(filename) == ".sql" {
			executed, err := migrationExecuted(ctx, db, gfile.Basename(filename))
			if err != nil {
				log.Printf("Error checking migration %s: %v\n", filename, err)
				continue
			}

			if !executed {
				err := executeSQLFile(ctx, db, filename)
				if err != nil {
					log.Printf("Error executing migration %s: %v\n", filename, err)
				} else {
					if err := recordMigration(ctx, db, gfile.Basename(filename)); err != nil {
						log.Printf("Error recording migration %s: %v\n", filename, err)
					}
				}
			} else {
				fmt.Printf("Migration %s has already been executed, skipping.\n", filename)
			}
		}
	}
}

func SQLiteMigration(ctx context.Context) error {
	database, err := sql.Open("sqlite3", "./overflows.db")
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	database.Close()

	db := g.DB("standaloneMode")
	if db == nil {
		return fmt.Errorf("database connection not initialized")
	}

	if err := createMigrationHistoryTable(ctx, db); err != nil {
		return fmt.Errorf("failed to create migration history table: %w", err)
	}

	executeMigrations(ctx, db, "manifest/database/SQLite/migrations")
	return nil
}
