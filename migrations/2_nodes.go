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
				`
				CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

				CREATE TABLE %[1]v (
					internal_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,

					name text,
					description text,
					score int,

					deposit_chain int,
					deposit_amount bigint,

					joined_at bigint,
					locked_until bigint,

					nebulas_using text[],

					contacts uuid,
					socials uuid
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