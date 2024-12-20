package migration

import (
	"database/sql"
	"dzrise.ru/internal/config"
	"dzrise.ru/internal/pkg/migraror"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
)

func main() {

	cnf := config.MustLoad()
	sqlStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cnf.Bd.User, cnf.Bd.Password, cnf.Bd.Host, cnf.Bd.Port, cnf.Bd.Name)

	db, err := sql.Open("postgres", sqlStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	m := migraror.New(sqlStr, cnf.Bd.Name, driver)
	m.Up()

}
