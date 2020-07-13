package main

import "fmt"

func CreateMaterializedViewQuery (tableName string) string {
	return fmt.Sprintf("CREATE MATERIALIZED VIEW %[1]v_materialized_view AS SELECT * FROM %[1]v;", tableName)
}

func DropMaterializedViewQuery (tableName string) string {
	return fmt.Sprintf("DROP MATERIALIZED VIEW %[1]v_materialized_view", tableName)
}
