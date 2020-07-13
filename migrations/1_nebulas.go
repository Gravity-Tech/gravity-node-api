package main

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/model"
	"github.com/go-pg/migrations"
)

func init () {
	tableName := model.DefaultExtendedDBTableNames.Nebulas

	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Printf("creating %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(
				`
				CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

				CREATE TABLE %[1]v (
					internal_id uuid DEFAULT uuid_generate_v4 (),
					address text PRIMARY KEY,
					name text,
					status int,
					description text,
					score int,
					target_chain int,
					subscription_fee text,

					nodes_using text[],
					regularity bigint
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