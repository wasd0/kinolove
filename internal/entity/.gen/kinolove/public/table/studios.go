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

var Studios = newStudiosTable("public", "studios", "")

type studiosTable struct {
	postgres.Table

	// Columns
	ID   postgres.ColumnInteger
	Name postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type StudiosTable struct {
	studiosTable

	EXCLUDED studiosTable
}

// AS creates new StudiosTable with assigned alias
func (a StudiosTable) AS(alias string) *StudiosTable {
	return newStudiosTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StudiosTable with assigned schema name
func (a StudiosTable) FromSchema(schemaName string) *StudiosTable {
	return newStudiosTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StudiosTable with assigned table prefix
func (a StudiosTable) WithPrefix(prefix string) *StudiosTable {
	return newStudiosTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StudiosTable with assigned table suffix
func (a StudiosTable) WithSuffix(suffix string) *StudiosTable {
	return newStudiosTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStudiosTable(schemaName, tableName, alias string) *StudiosTable {
	return &StudiosTable{
		studiosTable: newStudiosTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newStudiosTableImpl("", "excluded", ""),
	}
}

func newStudiosTableImpl(schemaName, tableName, alias string) studiosTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		NameColumn     = postgres.StringColumn("name")
		allColumns     = postgres.ColumnList{IDColumn, NameColumn}
		mutableColumns = postgres.ColumnList{NameColumn}
	)

	return studiosTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:   IDColumn,
		Name: NameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
