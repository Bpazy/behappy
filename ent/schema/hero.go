package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Hero holds the schema definition for the Hero entity.
type Hero struct {
	ent.Schema
}

// Fields of the Hero.
func (Hero) Fields() []ent.Field {
	return []ent.Field{
		field.Int("hero_id"),
		field.String("name"),
		field.String("localized_name"),
	}
}

func (Hero) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Hero.
func (Hero) Edges() []ent.Edge {
	return nil
}
