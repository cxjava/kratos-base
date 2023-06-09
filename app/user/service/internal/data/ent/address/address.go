// Code generated by ent, DO NOT EDIT.

package address

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the address type in the database.
	Label = "address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "street"
	// FieldJoinTime holds the string denoting the jointime field in the database.
	FieldJoinTime = "join_time"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// Table holds the table name of the address in the database.
	Table = "address"
)

// Columns holds all SQL columns for address fields.
var Columns = []string{
	FieldID,
	FieldStreet,
	FieldJoinTime,
	FieldHost,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Address queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStreet orders the results by the street field.
func ByStreet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStreet, opts...).ToFunc()
}

// ByJoinTime orders the results by the joinTime field.
func ByJoinTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJoinTime, opts...).ToFunc()
}

// ByHost orders the results by the host field.
func ByHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHost, opts...).ToFunc()
}
