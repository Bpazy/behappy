package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SubscriptionMatch holds the schema definition for the SubscriptionMatch entity.
type SubscriptionMatch struct {
	ent.Schema
}

// Fields of the SubscriptionMatch.
func (SubscriptionMatch) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("match_id"),
		field.String("player_id"),
		field.Int("player_slot"),
		field.Bool("radiant_win"),
		field.Int("duration"),
		field.Int("game_mode"),
		field.Int("lobby_type"),
		field.Int("hero_id"),
		field.Int("start_time"),
		field.Int("version"),
		field.Int("kills"),
		field.Int("deaths"),
		field.Int("assists"),
		field.Int("skill").Optional().Nillable(),
		field.Int("leaver_status"),
		field.Int("party_size"),
	}
}

// Mixin of the SubscriptionMatch.
func (SubscriptionMatch) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the SubscriptionMatch.
func (SubscriptionMatch) Edges() []ent.Edge {
	return nil
}
