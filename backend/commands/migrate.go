package commands

import (
	"log"
	"os"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

type MigrateCommand struct {
	Meta
}

func (cmd MigrateCommand) Help() string {
	return ""
}

func (cmd MigrateCommand) Synopsis() string {
	return "migrates database"
}

func (cmd MigrateCommand) Run(args []string) int {
	log.Printf("Migrating Postgres database")

	m, err := migrate.New("file://./migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Error migrating database: %s", err)
		return 1
	}

	if err = m.Up(); err != nil {
		log.Printf("Error migrating database: %s", err)
		return 1
	}

	log.Printf("Done")
	return 0
}
