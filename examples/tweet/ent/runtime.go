// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/like"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/schema"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	likeFields := schema.Like{}.Fields()
	_ = likeFields
	// likeDescLikedAt is the schema descriptor for liked_at field.
	likeDescLikedAt := likeFields[2].Descriptor()
	// like.DefaultLikedAt holds the default value on creation for the liked_at field.
	like.DefaultLikedAt = likeDescLikedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}