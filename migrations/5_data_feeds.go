package main

import (
	"fmt"
	"github.com/Gravity-Tech/gravity-node-api/migrations/common"
	"github.com/Gravity-Tech/gravity-node-api/model"
	"github.com/go-pg/migrations"
)

func init () {
	tableName := model.DefaultExtendedDBTableNames.Datafeeds

	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			fmt.Printf("creating %v table...\n", tableName)
			_, err := db.Exec(fmt.Sprintf(
				`
				CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

				CREATE TABLE %[1]v (
					internal_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
					data_feed_tag text,
					description text
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