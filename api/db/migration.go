
package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	migrations, err := ioutil.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	for _, migration := range migrations {
		if !strings.HasSuffix(migration.Name(), ".sql") {
			continue
		}

		content, err := ioutil.ReadFile(filepath.Join("migrations", migration.Name()))
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", migration.Name(), err)
		}

		log.Printf("Running migration: %s\n", migration.Name())
		
		// Split the content into individual statements
		statements := strings.Split(string(content), ";")
		
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			_, err = db.Exec(stmt)
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", migration.Name(), err)
			}
		}
		
		log.Printf("Successfully completed migration: %s\n", migration.Name())
	}

	return nil
}
