// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressColumns holds the columns for the "address" table.
	AddressColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "主键"},
		{Name: "street", Type: field.TypeInt32, Comment: "街道"},
		{Name: "join_time", Type: field.TypeTime, Comment: "加入时间"},
		{Name: "host", Type: field.TypeString, Comment: "户主"},
	}
	// AddressTable holds the schema information for the "address" table.
	AddressTable = &schema.Table{
		Name:       "address",
		Comment:    "Host Address",
		Columns:    AddressColumns,
		PrimaryKey: []*schema.Column{AddressColumns[0]},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "age", Type: field.TypeInt32, Comment: "年龄"},
		{Name: "name", Type: field.TypeString, Comment: "姓名"},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Comment:    "Comment that appears in both the schema and the generated code",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressTable,
		UserTable,
	}
)

func init() {
	AddressTable.Annotation = &entsql.Annotation{
		Table: "address",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}