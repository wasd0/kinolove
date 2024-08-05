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

var TitleStudio = newTitleStudioTable("public", "title_studio", "")

type titleStudioTable struct {
	postgres.Table

	// Columns
	TitleID  postgres.ColumnInteger
	StudioID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TitleStudioTable struct {
	titleStudioTable

	EXCLUDED titleStudioTable
}

// AS creates new TitleStudioTable with assigned alias
func (a TitleStudioTable) AS(alias string) *TitleStudioTable {
	return newTitleStudioTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TitleStudioTable with assigned schema name
func (a TitleStudioTable) FromSchema(schemaName string) *TitleStudioTable {
	return newTitleStudioTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TitleStudioTable with assigned table prefix
func (a TitleStudioTable) WithPrefix(prefix string) *TitleStudioTable {
	return newTitleStudioTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TitleStudioTable with assigned table suffix
func (a TitleStudioTable) WithSuffix(suffix string) *TitleStudioTable {
	return newTitleStudioTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTitleStudioTable(schemaName, tableName, alias string) *TitleStudioTable {
	return &TitleStudioTable{
		titleStudioTable: newTitleStudioTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newTitleStudioTableImpl("", "excluded", ""),
	}
}

func newTitleStudioTableImpl(schemaName, tableName, alias string) titleStudioTable {
	var (
		TitleIDColumn  = postgres.IntegerColumn("title_id")
		StudioIDColumn = postgres.IntegerColumn("studio_id")
		allColumns     = postgres.ColumnList{TitleIDColumn, StudioIDColumn}
		mutableColumns = postgres.ColumnList{TitleIDColumn, StudioIDColumn}
	)

	return titleStudioTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TitleID:  TitleIDColumn,
		StudioID: StudioIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
