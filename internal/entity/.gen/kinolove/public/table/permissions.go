//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Permissions = newPermissionsTable("public", "permissions", "")

type permissionsTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	Name        postgres.ColumnString
	Description postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PermissionsTable struct {
	permissionsTable

	EXCLUDED permissionsTable
}

// AS creates new PermissionsTable with assigned alias
func (a PermissionsTable) AS(alias string) *PermissionsTable {
	return newPermissionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PermissionsTable with assigned schema name
func (a PermissionsTable) FromSchema(schemaName string) *PermissionsTable {
	return newPermissionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PermissionsTable with assigned table prefix
func (a PermissionsTable) WithPrefix(prefix string) *PermissionsTable {
	return newPermissionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PermissionsTable with assigned table suffix
func (a PermissionsTable) WithSuffix(suffix string) *PermissionsTable {
	return newPermissionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPermissionsTable(schemaName, tableName, alias string) *PermissionsTable {
	return &PermissionsTable{
		permissionsTable: newPermissionsTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newPermissionsTableImpl("", "excluded", ""),
	}
}

func newPermissionsTableImpl(schemaName, tableName, alias string) permissionsTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		NameColumn        = postgres.StringColumn("name")
		DescriptionColumn = postgres.StringColumn("description")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, DescriptionColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, DescriptionColumn}
	)

	return permissionsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}