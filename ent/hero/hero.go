// Code generated by ent, DO NOT EDIT.

package hero

import (
	"time"
)

const (
	// Label holds the string label denoting the hero type in the database.
	Label = "hero"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldHeroID holds the string denoting the hero_id field in the database.
	FieldHeroID = "hero_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLocalizedName holds the string denoting the localized_name field in the database.
	FieldLocalizedName = "localized_name"
	// Table holds the table name of the hero in the database.
	Table = "heros"
)

// Columns holds all SQL columns for hero fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldHeroID,
	FieldName,
	FieldLocalizedName,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
