// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/tweet"
)

// Tweet is the model entity for the Tweet schema.
type Tweet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TweetQuery when eager-loading is set.
	Edges TweetEdges `json:"edges"`
}

// TweetEdges holds the relations/edges for other nodes in the graph.
type TweetEdges struct {
	// LikedUsers holds the value of the liked_users edge.
	LikedUsers []*User `json:"liked_users,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LikedUsersOrErr returns the LikedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) LikedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.LikedUsers, nil
	}
	return nil, &NotLoadedError{edge: "liked_users"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e TweetEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[1] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tweet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID:
			values[i] = new(sql.NullInt64)
		case tweet.FieldText:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tweet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tweet fields.
func (t *Tweet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tweet.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		}
	}
	return nil
}

// QueryLikedUsers queries the "liked_users" edge of the Tweet entity.
func (t *Tweet) QueryLikedUsers() *UserQuery {
	return (&TweetClient{config: t.config}).QueryLikedUsers(t)
}

// QueryLikes queries the "likes" edge of the Tweet entity.
func (t *Tweet) QueryLikes() *LikeQuery {
	return (&TweetClient{config: t.config}).QueryLikes(t)
}

// Update returns a builder for updating this Tweet.
// Note that you need to call Tweet.Unwrap() before calling this method if this Tweet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tweet) Update() *TweetUpdateOne {
	return (&TweetClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tweet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tweet) Unwrap() *Tweet {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tweet is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tweet) String() string {
	var builder strings.Builder
	builder.WriteString("Tweet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Tweets is a parsable slice of Tweet.
type Tweets []*Tweet

func (t Tweets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
