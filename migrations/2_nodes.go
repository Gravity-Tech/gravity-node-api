package main

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/migrations/common"
	"github.com/Gravity-Tech/gravity-node-api/model"
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
					address text,
					public_key text,
					name text,
					description text,
					score int,

					deposit_chain int,
					deposit_amount bigint,

					joined_at bigint,
					locked_until bigint,

					nebulas_using text[],
					
					PRIMARY KEY (public_key)
				);
				%v;
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