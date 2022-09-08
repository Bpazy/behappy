// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Bpazy/behappy/ent/hero"
	"github.com/Bpazy/behappy/ent/schema"
	"github.com/Bpazy/behappy/ent/subscription"
	"github.com/Bpazy/behappy/ent/subscriptionmatch"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	heroMixin := schema.Hero{}.Mixin()
	heroMixinFields0 := heroMixin[0].Fields()
	_ = heroMixinFields0
	heroFields := schema.Hero{}.Fields()
	_ = heroFields
	// heroDescCreateTime is the schema descriptor for create_time field.
	heroDescCreateTime := heroMixinFields0[0].Descriptor()
	// hero.DefaultCreateTime holds the default value on creation for the create_time field.
	hero.DefaultCreateTime = heroDescCreateTime.Default.(func() time.Time)
	// heroDescUpdateTime is the schema descriptor for update_time field.
	heroDescUpdateTime := heroMixinFields0[1].Descriptor()
	// hero.DefaultUpdateTime holds the default value on creation for the update_time field.
	hero.DefaultUpdateTime = heroDescUpdateTime.Default.(func() time.Time)
	// hero.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	hero.UpdateDefaultUpdateTime = heroDescUpdateTime.UpdateDefault.(func() time.Time)
	subscriptionMixin := schema.Subscription{}.Mixin()
	subscriptionMixinFields0 := subscriptionMixin[0].Fields()
	_ = subscriptionMixinFields0
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescCreateTime is the schema descriptor for create_time field.
	subscriptionDescCreateTime := subscriptionMixinFields0[0].Descriptor()
	// subscription.DefaultCreateTime holds the default value on creation for the create_time field.
	subscription.DefaultCreateTime = subscriptionDescCreateTime.Default.(func() time.Time)
	// subscriptionDescUpdateTime is the schema descriptor for update_time field.
	subscriptionDescUpdateTime := subscriptionMixinFields0[1].Descriptor()
	// subscription.DefaultUpdateTime holds the default value on creation for the update_time field.
	subscription.DefaultUpdateTime = subscriptionDescUpdateTime.Default.(func() time.Time)
	// subscription.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	subscription.UpdateDefaultUpdateTime = subscriptionDescUpdateTime.UpdateDefault.(func() time.Time)
	subscriptionmatchMixin := schema.SubscriptionMatch{}.Mixin()
	subscriptionmatchMixinFields0 := subscriptionmatchMixin[0].Fields()
	_ = subscriptionmatchMixinFields0
	subscriptionmatchFields := schema.SubscriptionMatch{}.Fields()
	_ = subscriptionmatchFields
	// subscriptionmatchDescCreateTime is the schema descriptor for create_time field.
	subscriptionmatchDescCreateTime := subscriptionmatchMixinFields0[0].Descriptor()
	// subscriptionmatch.DefaultCreateTime holds the default value on creation for the create_time field.
	subscriptionmatch.DefaultCreateTime = subscriptionmatchDescCreateTime.Default.(func() time.Time)
	// subscriptionmatchDescUpdateTime is the schema descriptor for update_time field.
	subscriptionmatchDescUpdateTime := subscriptionmatchMixinFields0[1].Descriptor()
	// subscriptionmatch.DefaultUpdateTime holds the default value on creation for the update_time field.
	subscriptionmatch.DefaultUpdateTime = subscriptionmatchDescUpdateTime.Default.(func() time.Time)
	// subscriptionmatch.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	subscriptionmatch.UpdateDefaultUpdateTime = subscriptionmatchDescUpdateTime.UpdateDefault.(func() time.Time)
}
