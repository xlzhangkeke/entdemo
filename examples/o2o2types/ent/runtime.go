// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/xlzhangkeke/entdemo/examples/o2o2types/ent/schema"
	"github.com/xlzhangkeke/entdemo/examples/o2o2types/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
}
