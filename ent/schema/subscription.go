package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

const (
	fieldGroupId = "group_id"
	fieldSteamId = "steam_id"
	fieldAlias   = "alias"
)

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.Int(fieldGroupId),
		field.String(fieldSteamId),
		field.String(fieldAlias),
	}
}

func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Subscription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(fieldGroupId),
		index.Fields(fieldSteamId),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return nil
}
