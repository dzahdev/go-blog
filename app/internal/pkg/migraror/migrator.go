package migraror

import (
	"dzrise.ru/internal/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"log"
	"log/slog"
)

type Migrator struct {
	m   *migrate.Migrate
	log *slog.Logger
}

func New(cnf string, dbName string, driver database.Driver) *Migrator {
	log := logger.SetupLogger("development")
	m, err := migrate.NewWithDatabaseInstance(
		cnf,
		dbName,
		driver)
	if err != nil {
		log.Error(err.Error())
	}
	return &Migrator{
		m:   m,
		log: log,
	}
}

func (m *Migrator) Up() {
	err := m.m.Up()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (m *Migrator) Down() {
	err := m.m.Down()
	if err != nil {
		log.Fatal(err.Error())
	}
}
