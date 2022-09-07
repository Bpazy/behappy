package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const (
	fieldHeroId        = "hero_id"
	fieldName          = "name"
	fieldLocalizedName = "localized_name"
)

// Hero holds the schema definition for the Hero entity.
type Hero struct {
	ent.Schema
}

// Fields of the Hero.
func (Hero) Fields() []ent.Field {
	return []ent.Field{
		field.Int(fieldHeroId),
		field.String(fieldName),
		field.String(fieldLocalizedName),
	}
}

func (Hero) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Hero) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(fieldHeroId),
	}
}

// Edges of the Hero.
func (Hero) Edges() []ent.Edge {
	return nil
}
