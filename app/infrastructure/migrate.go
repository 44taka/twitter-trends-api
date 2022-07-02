package infrastructure

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate(c *Config) error {
	// TODO:migrateのバージョンをv4に変える
	m, err := migrate.New(c.Migration.FILE_URI, c.DB.URI)
	if err != nil {
		fmt.Println("err 1")
		return err
	}
	err = m.Up()
	if err != nil {
		fmt.Println("err 2")
		if err != migrate.ErrNoChange && err != migrate.ErrNilVersion {
			fmt.Println("err 3")
			return err
		}
	}
	return nil
}
