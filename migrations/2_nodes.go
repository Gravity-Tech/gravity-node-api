package main

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"github.com/go-pg/migrations"
)

func init () {
	tableName := model.DefaultExtendedDBTableNames.Nodes

	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Printf("creating %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(
				`CREATE TABLE %[1]v (
					foo text,
					bar int
				);
				`, tableName))
			return err
			//return nil
		},
		func(db migrations.DB) error {
			fmt.Printf("dropping %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(`DROP TABLE %v`, tableName))
			return err
		},
	)
}