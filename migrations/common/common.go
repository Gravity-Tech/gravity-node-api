package common

import "fmt"

const (
	materializedViewPostfix = "_materialized_view"
)

func CreateMaterializedViewQuery (tableName string) string {
	return fmt.Sprintf("CREATE MATERIALIZED VIEW %v AS SELECT * FROM %v;", tableName + materializedViewPostfix, tableName)
}

func DropMaterializedViewQuery (tableName string) string {
	return fmt.Sprintf("DROP MATERIALIZED VIEW %[1]v;", tableName + materializedViewPostfix)
}

func UpdateMaterializedViewQuery (tableName string) string {
	return fmt.Sprintf("REFRESH MATERIALIZED VIEW %[1]v;", tableName + materializedViewPostfix)
}
