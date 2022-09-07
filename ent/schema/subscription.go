package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id"),
		field.String("steam_id"),
		field.String("alias"),
	}
}

func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return nil
}
