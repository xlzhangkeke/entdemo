// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeLikedTweets holds the string denoting the liked_tweets edge name in mutations.
	EdgeLikedTweets = "liked_tweets"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// LikedTweetsTable is the table that holds the liked_tweets relation/edge. The primary key declared below.
	LikedTweetsTable = "likes"
	// LikedTweetsInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	LikedTweetsInverseTable = "tweets"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// LikedTweetsPrimaryKey and LikedTweetsColumn2 are the table columns denoting the
	// primary key for the liked_tweets relation (M2M).
	LikedTweetsPrimaryKey = []string{"user_id", "tweet_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
