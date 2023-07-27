// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CollectsColumns holds the columns for the "collects" table.
	CollectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CollectsTable holds the schema information for the "collects" table.
	CollectsTable = &schema.Table{
		Name:       "collects",
		Columns:    CollectsColumns,
		PrimaryKey: []*schema.Column{CollectsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CollectsTable,
	}
)

func init() {
}
