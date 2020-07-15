package main

import (
	"fmt"
	"github.com/Gravity-Hub-Org/gravity-node-api-mockup/v2/migrations/common"
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
				
				%[2]v;
				`, tableName, common.CreateMaterializedViewQuery(tableName)))
			return err
		},
		func(db migrations.DB) error {
			fmt.Printf("dropping %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(`%v; DROP TABLE %v;`, common.DropMaterializedViewQuery(tableName), tableName))

			return err
		},
	)
}