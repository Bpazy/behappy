// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HerosColumns holds the columns for the "heros" table.
	HerosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "hero_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "localized_name", Type: field.TypeString},
	}
	// HerosTable holds the schema information for the "heros" table.
	HerosTable = &schema.Table{
		Name:       "heros",
		Columns:    HerosColumns,
		PrimaryKey: []*schema.Column{HerosColumns[0]},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "group_id", Type: field.TypeInt},
		{Name: "steam_id", Type: field.TypeString},
		{Name: "alias", Type: field.TypeString},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HerosTable,
		SubscriptionsTable,
	}
)

func init() {
}
