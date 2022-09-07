// Code generated by ent, DO NOT EDIT.

package hero

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Bpazy/behappy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// HeroID applies equality check predicate on the "hero_id" field. It's identical to HeroIDEQ.
func HeroID(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeroID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// LocalizedName applies equality check predicate on the "localized_name" field. It's identical to LocalizedNameEQ.
func LocalizedName(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalizedName), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// HeroIDEQ applies the EQ predicate on the "hero_id" field.
func HeroIDEQ(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeroID), v))
	})
}

// HeroIDNEQ applies the NEQ predicate on the "hero_id" field.
func HeroIDNEQ(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeroID), v))
	})
}

// HeroIDIn applies the In predicate on the "hero_id" field.
func HeroIDIn(vs ...int) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHeroID), v...))
	})
}

// HeroIDNotIn applies the NotIn predicate on the "hero_id" field.
func HeroIDNotIn(vs ...int) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHeroID), v...))
	})
}

// HeroIDGT applies the GT predicate on the "hero_id" field.
func HeroIDGT(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHeroID), v))
	})
}

// HeroIDGTE applies the GTE predicate on the "hero_id" field.
func HeroIDGTE(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHeroID), v))
	})
}

// HeroIDLT applies the LT predicate on the "hero_id" field.
func HeroIDLT(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHeroID), v))
	})
}

// HeroIDLTE applies the LTE predicate on the "hero_id" field.
func HeroIDLTE(v int) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHeroID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// LocalizedNameEQ applies the EQ predicate on the "localized_name" field.
func LocalizedNameEQ(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameNEQ applies the NEQ predicate on the "localized_name" field.
func LocalizedNameNEQ(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameIn applies the In predicate on the "localized_name" field.
func LocalizedNameIn(vs ...string) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocalizedName), v...))
	})
}

// LocalizedNameNotIn applies the NotIn predicate on the "localized_name" field.
func LocalizedNameNotIn(vs ...string) predicate.Hero {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocalizedName), v...))
	})
}

// LocalizedNameGT applies the GT predicate on the "localized_name" field.
func LocalizedNameGT(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameGTE applies the GTE predicate on the "localized_name" field.
func LocalizedNameGTE(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameLT applies the LT predicate on the "localized_name" field.
func LocalizedNameLT(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameLTE applies the LTE predicate on the "localized_name" field.
func LocalizedNameLTE(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameContains applies the Contains predicate on the "localized_name" field.
func LocalizedNameContains(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameHasPrefix applies the HasPrefix predicate on the "localized_name" field.
func LocalizedNameHasPrefix(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameHasSuffix applies the HasSuffix predicate on the "localized_name" field.
func LocalizedNameHasSuffix(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameEqualFold applies the EqualFold predicate on the "localized_name" field.
func LocalizedNameEqualFold(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocalizedName), v))
	})
}

// LocalizedNameContainsFold applies the ContainsFold predicate on the "localized_name" field.
func LocalizedNameContainsFold(v string) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocalizedName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Hero) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Hero) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Hero) predicate.Hero {
	return predicate.Hero(func(s *sql.Selector) {
		p(s.Not())
	})
}
