// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Bpazy/behappy/ent/predicate"
	"github.com/Bpazy/behappy/ent/subscriptionmatch"
)

// SubscriptionMatchUpdate is the builder for updating SubscriptionMatch entities.
type SubscriptionMatchUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionMatchMutation
}

// Where appends a list predicates to the SubscriptionMatchUpdate builder.
func (smu *SubscriptionMatchUpdate) Where(ps ...predicate.SubscriptionMatch) *SubscriptionMatchUpdate {
	smu.mutation.Where(ps...)
	return smu
}

// SetUpdateTime sets the "update_time" field.
func (smu *SubscriptionMatchUpdate) SetUpdateTime(t time.Time) *SubscriptionMatchUpdate {
	smu.mutation.SetUpdateTime(t)
	return smu
}

// SetMatchID sets the "match_id" field.
func (smu *SubscriptionMatchUpdate) SetMatchID(i int64) *SubscriptionMatchUpdate {
	smu.mutation.ResetMatchID()
	smu.mutation.SetMatchID(i)
	return smu
}

// AddMatchID adds i to the "match_id" field.
func (smu *SubscriptionMatchUpdate) AddMatchID(i int64) *SubscriptionMatchUpdate {
	smu.mutation.AddMatchID(i)
	return smu
}

// SetPlayerID sets the "player_id" field.
func (smu *SubscriptionMatchUpdate) SetPlayerID(s string) *SubscriptionMatchUpdate {
	smu.mutation.SetPlayerID(s)
	return smu
}

// SetPlayerSlot sets the "player_slot" field.
func (smu *SubscriptionMatchUpdate) SetPlayerSlot(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetPlayerSlot()
	smu.mutation.SetPlayerSlot(i)
	return smu
}

// AddPlayerSlot adds i to the "player_slot" field.
func (smu *SubscriptionMatchUpdate) AddPlayerSlot(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddPlayerSlot(i)
	return smu
}

// SetRadiantWin sets the "radiant_win" field.
func (smu *SubscriptionMatchUpdate) SetRadiantWin(b bool) *SubscriptionMatchUpdate {
	smu.mutation.SetRadiantWin(b)
	return smu
}

// SetDuration sets the "duration" field.
func (smu *SubscriptionMatchUpdate) SetDuration(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetDuration()
	smu.mutation.SetDuration(i)
	return smu
}

// AddDuration adds i to the "duration" field.
func (smu *SubscriptionMatchUpdate) AddDuration(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddDuration(i)
	return smu
}

// SetGameMode sets the "game_mode" field.
func (smu *SubscriptionMatchUpdate) SetGameMode(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetGameMode()
	smu.mutation.SetGameMode(i)
	return smu
}

// AddGameMode adds i to the "game_mode" field.
func (smu *SubscriptionMatchUpdate) AddGameMode(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddGameMode(i)
	return smu
}

// SetLobbyType sets the "lobby_type" field.
func (smu *SubscriptionMatchUpdate) SetLobbyType(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetLobbyType()
	smu.mutation.SetLobbyType(i)
	return smu
}

// AddLobbyType adds i to the "lobby_type" field.
func (smu *SubscriptionMatchUpdate) AddLobbyType(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddLobbyType(i)
	return smu
}

// SetHeroID sets the "hero_id" field.
func (smu *SubscriptionMatchUpdate) SetHeroID(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetHeroID()
	smu.mutation.SetHeroID(i)
	return smu
}

// AddHeroID adds i to the "hero_id" field.
func (smu *SubscriptionMatchUpdate) AddHeroID(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddHeroID(i)
	return smu
}

// SetStartTime sets the "start_time" field.
func (smu *SubscriptionMatchUpdate) SetStartTime(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetStartTime()
	smu.mutation.SetStartTime(i)
	return smu
}

// AddStartTime adds i to the "start_time" field.
func (smu *SubscriptionMatchUpdate) AddStartTime(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddStartTime(i)
	return smu
}

// SetVersion sets the "version" field.
func (smu *SubscriptionMatchUpdate) SetVersion(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetVersion()
	smu.mutation.SetVersion(i)
	return smu
}

// AddVersion adds i to the "version" field.
func (smu *SubscriptionMatchUpdate) AddVersion(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddVersion(i)
	return smu
}

// SetKills sets the "kills" field.
func (smu *SubscriptionMatchUpdate) SetKills(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetKills()
	smu.mutation.SetKills(i)
	return smu
}

// AddKills adds i to the "kills" field.
func (smu *SubscriptionMatchUpdate) AddKills(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddKills(i)
	return smu
}

// SetDeaths sets the "deaths" field.
func (smu *SubscriptionMatchUpdate) SetDeaths(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetDeaths()
	smu.mutation.SetDeaths(i)
	return smu
}

// AddDeaths adds i to the "deaths" field.
func (smu *SubscriptionMatchUpdate) AddDeaths(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddDeaths(i)
	return smu
}

// SetAssists sets the "assists" field.
func (smu *SubscriptionMatchUpdate) SetAssists(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetAssists()
	smu.mutation.SetAssists(i)
	return smu
}

// AddAssists adds i to the "assists" field.
func (smu *SubscriptionMatchUpdate) AddAssists(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddAssists(i)
	return smu
}

// SetSkill sets the "skill" field.
func (smu *SubscriptionMatchUpdate) SetSkill(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetSkill()
	smu.mutation.SetSkill(i)
	return smu
}

// SetNillableSkill sets the "skill" field if the given value is not nil.
func (smu *SubscriptionMatchUpdate) SetNillableSkill(i *int) *SubscriptionMatchUpdate {
	if i != nil {
		smu.SetSkill(*i)
	}
	return smu
}

// AddSkill adds i to the "skill" field.
func (smu *SubscriptionMatchUpdate) AddSkill(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddSkill(i)
	return smu
}

// ClearSkill clears the value of the "skill" field.
func (smu *SubscriptionMatchUpdate) ClearSkill() *SubscriptionMatchUpdate {
	smu.mutation.ClearSkill()
	return smu
}

// SetLeaverStatus sets the "leaver_status" field.
func (smu *SubscriptionMatchUpdate) SetLeaverStatus(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetLeaverStatus()
	smu.mutation.SetLeaverStatus(i)
	return smu
}

// AddLeaverStatus adds i to the "leaver_status" field.
func (smu *SubscriptionMatchUpdate) AddLeaverStatus(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddLeaverStatus(i)
	return smu
}

// SetPartySize sets the "party_size" field.
func (smu *SubscriptionMatchUpdate) SetPartySize(i int) *SubscriptionMatchUpdate {
	smu.mutation.ResetPartySize()
	smu.mutation.SetPartySize(i)
	return smu
}

// AddPartySize adds i to the "party_size" field.
func (smu *SubscriptionMatchUpdate) AddPartySize(i int) *SubscriptionMatchUpdate {
	smu.mutation.AddPartySize(i)
	return smu
}

// Mutation returns the SubscriptionMatchMutation object of the builder.
func (smu *SubscriptionMatchUpdate) Mutation() *SubscriptionMatchMutation {
	return smu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *SubscriptionMatchUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	smu.defaults()
	if len(smu.hooks) == 0 {
		affected, err = smu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMatchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smu.mutation = mutation
			affected, err = smu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smu.hooks) - 1; i >= 0; i-- {
			if smu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (smu *SubscriptionMatchUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *SubscriptionMatchUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *SubscriptionMatchUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smu *SubscriptionMatchUpdate) defaults() {
	if _, ok := smu.mutation.UpdateTime(); !ok {
		v := subscriptionmatch.UpdateDefaultUpdateTime()
		smu.mutation.SetUpdateTime(v)
	}
}

func (smu *SubscriptionMatchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscriptionmatch.Table,
			Columns: subscriptionmatch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscriptionmatch.FieldID,
			},
		},
	}
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscriptionmatch.FieldUpdateTime,
		})
	}
	if value, ok := smu.mutation.MatchID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: subscriptionmatch.FieldMatchID,
		})
	}
	if value, ok := smu.mutation.AddedMatchID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: subscriptionmatch.FieldMatchID,
		})
	}
	if value, ok := smu.mutation.PlayerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptionmatch.FieldPlayerID,
		})
	}
	if value, ok := smu.mutation.PlayerSlot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPlayerSlot,
		})
	}
	if value, ok := smu.mutation.AddedPlayerSlot(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPlayerSlot,
		})
	}
	if value, ok := smu.mutation.RadiantWin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subscriptionmatch.FieldRadiantWin,
		})
	}
	if value, ok := smu.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDuration,
		})
	}
	if value, ok := smu.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDuration,
		})
	}
	if value, ok := smu.mutation.GameMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldGameMode,
		})
	}
	if value, ok := smu.mutation.AddedGameMode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldGameMode,
		})
	}
	if value, ok := smu.mutation.LobbyType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLobbyType,
		})
	}
	if value, ok := smu.mutation.AddedLobbyType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLobbyType,
		})
	}
	if value, ok := smu.mutation.HeroID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldHeroID,
		})
	}
	if value, ok := smu.mutation.AddedHeroID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldHeroID,
		})
	}
	if value, ok := smu.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldStartTime,
		})
	}
	if value, ok := smu.mutation.AddedStartTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldStartTime,
		})
	}
	if value, ok := smu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldVersion,
		})
	}
	if value, ok := smu.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldVersion,
		})
	}
	if value, ok := smu.mutation.Kills(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldKills,
		})
	}
	if value, ok := smu.mutation.AddedKills(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldKills,
		})
	}
	if value, ok := smu.mutation.Deaths(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDeaths,
		})
	}
	if value, ok := smu.mutation.AddedDeaths(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDeaths,
		})
	}
	if value, ok := smu.mutation.Assists(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldAssists,
		})
	}
	if value, ok := smu.mutation.AddedAssists(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldAssists,
		})
	}
	if value, ok := smu.mutation.Skill(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldSkill,
		})
	}
	if value, ok := smu.mutation.AddedSkill(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldSkill,
		})
	}
	if smu.mutation.SkillCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: subscriptionmatch.FieldSkill,
		})
	}
	if value, ok := smu.mutation.LeaverStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLeaverStatus,
		})
	}
	if value, ok := smu.mutation.AddedLeaverStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLeaverStatus,
		})
	}
	if value, ok := smu.mutation.PartySize(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPartySize,
		})
	}
	if value, ok := smu.mutation.AddedPartySize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPartySize,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscriptionmatch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SubscriptionMatchUpdateOne is the builder for updating a single SubscriptionMatch entity.
type SubscriptionMatchUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionMatchMutation
}

// SetUpdateTime sets the "update_time" field.
func (smuo *SubscriptionMatchUpdateOne) SetUpdateTime(t time.Time) *SubscriptionMatchUpdateOne {
	smuo.mutation.SetUpdateTime(t)
	return smuo
}

// SetMatchID sets the "match_id" field.
func (smuo *SubscriptionMatchUpdateOne) SetMatchID(i int64) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetMatchID()
	smuo.mutation.SetMatchID(i)
	return smuo
}

// AddMatchID adds i to the "match_id" field.
func (smuo *SubscriptionMatchUpdateOne) AddMatchID(i int64) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddMatchID(i)
	return smuo
}

// SetPlayerID sets the "player_id" field.
func (smuo *SubscriptionMatchUpdateOne) SetPlayerID(s string) *SubscriptionMatchUpdateOne {
	smuo.mutation.SetPlayerID(s)
	return smuo
}

// SetPlayerSlot sets the "player_slot" field.
func (smuo *SubscriptionMatchUpdateOne) SetPlayerSlot(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetPlayerSlot()
	smuo.mutation.SetPlayerSlot(i)
	return smuo
}

// AddPlayerSlot adds i to the "player_slot" field.
func (smuo *SubscriptionMatchUpdateOne) AddPlayerSlot(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddPlayerSlot(i)
	return smuo
}

// SetRadiantWin sets the "radiant_win" field.
func (smuo *SubscriptionMatchUpdateOne) SetRadiantWin(b bool) *SubscriptionMatchUpdateOne {
	smuo.mutation.SetRadiantWin(b)
	return smuo
}

// SetDuration sets the "duration" field.
func (smuo *SubscriptionMatchUpdateOne) SetDuration(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetDuration()
	smuo.mutation.SetDuration(i)
	return smuo
}

// AddDuration adds i to the "duration" field.
func (smuo *SubscriptionMatchUpdateOne) AddDuration(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddDuration(i)
	return smuo
}

// SetGameMode sets the "game_mode" field.
func (smuo *SubscriptionMatchUpdateOne) SetGameMode(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetGameMode()
	smuo.mutation.SetGameMode(i)
	return smuo
}

// AddGameMode adds i to the "game_mode" field.
func (smuo *SubscriptionMatchUpdateOne) AddGameMode(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddGameMode(i)
	return smuo
}

// SetLobbyType sets the "lobby_type" field.
func (smuo *SubscriptionMatchUpdateOne) SetLobbyType(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetLobbyType()
	smuo.mutation.SetLobbyType(i)
	return smuo
}

// AddLobbyType adds i to the "lobby_type" field.
func (smuo *SubscriptionMatchUpdateOne) AddLobbyType(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddLobbyType(i)
	return smuo
}

// SetHeroID sets the "hero_id" field.
func (smuo *SubscriptionMatchUpdateOne) SetHeroID(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetHeroID()
	smuo.mutation.SetHeroID(i)
	return smuo
}

// AddHeroID adds i to the "hero_id" field.
func (smuo *SubscriptionMatchUpdateOne) AddHeroID(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddHeroID(i)
	return smuo
}

// SetStartTime sets the "start_time" field.
func (smuo *SubscriptionMatchUpdateOne) SetStartTime(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetStartTime()
	smuo.mutation.SetStartTime(i)
	return smuo
}

// AddStartTime adds i to the "start_time" field.
func (smuo *SubscriptionMatchUpdateOne) AddStartTime(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddStartTime(i)
	return smuo
}

// SetVersion sets the "version" field.
func (smuo *SubscriptionMatchUpdateOne) SetVersion(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetVersion()
	smuo.mutation.SetVersion(i)
	return smuo
}

// AddVersion adds i to the "version" field.
func (smuo *SubscriptionMatchUpdateOne) AddVersion(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddVersion(i)
	return smuo
}

// SetKills sets the "kills" field.
func (smuo *SubscriptionMatchUpdateOne) SetKills(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetKills()
	smuo.mutation.SetKills(i)
	return smuo
}

// AddKills adds i to the "kills" field.
func (smuo *SubscriptionMatchUpdateOne) AddKills(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddKills(i)
	return smuo
}

// SetDeaths sets the "deaths" field.
func (smuo *SubscriptionMatchUpdateOne) SetDeaths(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetDeaths()
	smuo.mutation.SetDeaths(i)
	return smuo
}

// AddDeaths adds i to the "deaths" field.
func (smuo *SubscriptionMatchUpdateOne) AddDeaths(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddDeaths(i)
	return smuo
}

// SetAssists sets the "assists" field.
func (smuo *SubscriptionMatchUpdateOne) SetAssists(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetAssists()
	smuo.mutation.SetAssists(i)
	return smuo
}

// AddAssists adds i to the "assists" field.
func (smuo *SubscriptionMatchUpdateOne) AddAssists(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddAssists(i)
	return smuo
}

// SetSkill sets the "skill" field.
func (smuo *SubscriptionMatchUpdateOne) SetSkill(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetSkill()
	smuo.mutation.SetSkill(i)
	return smuo
}

// SetNillableSkill sets the "skill" field if the given value is not nil.
func (smuo *SubscriptionMatchUpdateOne) SetNillableSkill(i *int) *SubscriptionMatchUpdateOne {
	if i != nil {
		smuo.SetSkill(*i)
	}
	return smuo
}

// AddSkill adds i to the "skill" field.
func (smuo *SubscriptionMatchUpdateOne) AddSkill(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddSkill(i)
	return smuo
}

// ClearSkill clears the value of the "skill" field.
func (smuo *SubscriptionMatchUpdateOne) ClearSkill() *SubscriptionMatchUpdateOne {
	smuo.mutation.ClearSkill()
	return smuo
}

// SetLeaverStatus sets the "leaver_status" field.
func (smuo *SubscriptionMatchUpdateOne) SetLeaverStatus(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetLeaverStatus()
	smuo.mutation.SetLeaverStatus(i)
	return smuo
}

// AddLeaverStatus adds i to the "leaver_status" field.
func (smuo *SubscriptionMatchUpdateOne) AddLeaverStatus(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddLeaverStatus(i)
	return smuo
}

// SetPartySize sets the "party_size" field.
func (smuo *SubscriptionMatchUpdateOne) SetPartySize(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.ResetPartySize()
	smuo.mutation.SetPartySize(i)
	return smuo
}

// AddPartySize adds i to the "party_size" field.
func (smuo *SubscriptionMatchUpdateOne) AddPartySize(i int) *SubscriptionMatchUpdateOne {
	smuo.mutation.AddPartySize(i)
	return smuo
}

// Mutation returns the SubscriptionMatchMutation object of the builder.
func (smuo *SubscriptionMatchUpdateOne) Mutation() *SubscriptionMatchMutation {
	return smuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *SubscriptionMatchUpdateOne) Select(field string, fields ...string) *SubscriptionMatchUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated SubscriptionMatch entity.
func (smuo *SubscriptionMatchUpdateOne) Save(ctx context.Context) (*SubscriptionMatch, error) {
	var (
		err  error
		node *SubscriptionMatch
	)
	smuo.defaults()
	if len(smuo.hooks) == 0 {
		node, err = smuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionMatchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			smuo.mutation = mutation
			node, err = smuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smuo.hooks) - 1; i >= 0; i-- {
			if smuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = smuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, smuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SubscriptionMatch)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubscriptionMatchMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *SubscriptionMatchUpdateOne) SaveX(ctx context.Context) *SubscriptionMatch {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *SubscriptionMatchUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *SubscriptionMatchUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smuo *SubscriptionMatchUpdateOne) defaults() {
	if _, ok := smuo.mutation.UpdateTime(); !ok {
		v := subscriptionmatch.UpdateDefaultUpdateTime()
		smuo.mutation.SetUpdateTime(v)
	}
}

func (smuo *SubscriptionMatchUpdateOne) sqlSave(ctx context.Context) (_node *SubscriptionMatch, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscriptionmatch.Table,
			Columns: subscriptionmatch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscriptionmatch.FieldID,
			},
		},
	}
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubscriptionMatch.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscriptionmatch.FieldID)
		for _, f := range fields {
			if !subscriptionmatch.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscriptionmatch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscriptionmatch.FieldUpdateTime,
		})
	}
	if value, ok := smuo.mutation.MatchID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: subscriptionmatch.FieldMatchID,
		})
	}
	if value, ok := smuo.mutation.AddedMatchID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: subscriptionmatch.FieldMatchID,
		})
	}
	if value, ok := smuo.mutation.PlayerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptionmatch.FieldPlayerID,
		})
	}
	if value, ok := smuo.mutation.PlayerSlot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPlayerSlot,
		})
	}
	if value, ok := smuo.mutation.AddedPlayerSlot(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPlayerSlot,
		})
	}
	if value, ok := smuo.mutation.RadiantWin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subscriptionmatch.FieldRadiantWin,
		})
	}
	if value, ok := smuo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDuration,
		})
	}
	if value, ok := smuo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDuration,
		})
	}
	if value, ok := smuo.mutation.GameMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldGameMode,
		})
	}
	if value, ok := smuo.mutation.AddedGameMode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldGameMode,
		})
	}
	if value, ok := smuo.mutation.LobbyType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLobbyType,
		})
	}
	if value, ok := smuo.mutation.AddedLobbyType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLobbyType,
		})
	}
	if value, ok := smuo.mutation.HeroID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldHeroID,
		})
	}
	if value, ok := smuo.mutation.AddedHeroID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldHeroID,
		})
	}
	if value, ok := smuo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldStartTime,
		})
	}
	if value, ok := smuo.mutation.AddedStartTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldStartTime,
		})
	}
	if value, ok := smuo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldVersion,
		})
	}
	if value, ok := smuo.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldVersion,
		})
	}
	if value, ok := smuo.mutation.Kills(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldKills,
		})
	}
	if value, ok := smuo.mutation.AddedKills(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldKills,
		})
	}
	if value, ok := smuo.mutation.Deaths(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDeaths,
		})
	}
	if value, ok := smuo.mutation.AddedDeaths(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldDeaths,
		})
	}
	if value, ok := smuo.mutation.Assists(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldAssists,
		})
	}
	if value, ok := smuo.mutation.AddedAssists(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldAssists,
		})
	}
	if value, ok := smuo.mutation.Skill(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldSkill,
		})
	}
	if value, ok := smuo.mutation.AddedSkill(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldSkill,
		})
	}
	if smuo.mutation.SkillCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: subscriptionmatch.FieldSkill,
		})
	}
	if value, ok := smuo.mutation.LeaverStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLeaverStatus,
		})
	}
	if value, ok := smuo.mutation.AddedLeaverStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldLeaverStatus,
		})
	}
	if value, ok := smuo.mutation.PartySize(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPartySize,
		})
	}
	if value, ok := smuo.mutation.AddedPartySize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subscriptionmatch.FieldPartySize,
		})
	}
	_node = &SubscriptionMatch{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscriptionmatch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}