// Code generated by ent, DO NOT EDIT.

package account

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeMutes holds the string denoting the mutes edge name in mutations.
	EdgeMutes = "mutes"
	// EdgeRequests holds the string denoting the requests edge name in mutations.
	EdgeRequests = "requests"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeSession holds the string denoting the session edge name in mutations.
	EdgeSession = "session"
	// EdgeFriendships holds the string denoting the friendships edge name in mutations.
	EdgeFriendships = "friendships"
	// EdgeRequestTargets holds the string denoting the requesttargets edge name in mutations.
	EdgeRequestTargets = "requestTargets"
	// MuteFieldID holds the string denoting the ID field of the Mute.
	MuteFieldID = "id"
	// LikeFieldID holds the string denoting the ID field of the Like.
	LikeFieldID = "id"
	// SessionFieldID holds the string denoting the ID field of the Session.
	SessionFieldID = "id"
	// FriendFieldID holds the string denoting the ID field of the Friend.
	FriendFieldID = "id"
	// RequestFieldID holds the string denoting the ID field of the Request.
	RequestFieldID = "id"
	// Table holds the table name of the account in the database.
	Table = "Accounts"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "Posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "Posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "account_id"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "Comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "Comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "account_id"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "Friends"
	// MutesTable is the table that holds the mutes relation/edge.
	MutesTable = "Mutes"
	// MutesInverseTable is the table name for the Mute entity.
	// It exists in this package in order to avoid circular dependency with the "mute" package.
	MutesInverseTable = "Mutes"
	// MutesColumn is the table column denoting the mutes relation/edge.
	MutesColumn = "account_mutes"
	// RequestsTable is the table that holds the requests relation/edge. The primary key declared below.
	RequestsTable = "Requests"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "Friends"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "Friends"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "account_likes"
	// SessionTable is the table that holds the session relation/edge.
	SessionTable = "Sessions"
	// SessionInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionInverseTable = "Sessions"
	// SessionColumn is the table column denoting the session relation/edge.
	SessionColumn = "account_id"
	// FriendshipsTable is the table that holds the friendships relation/edge.
	FriendshipsTable = "Friends"
	// FriendshipsInverseTable is the table name for the Friend entity.
	// It exists in this package in order to avoid circular dependency with the "friend" package.
	FriendshipsInverseTable = "Friends"
	// FriendshipsColumn is the table column denoting the friendships relation/edge.
	FriendshipsColumn = "account_id"
	// RequestTargetsTable is the table that holds the requestTargets relation/edge.
	RequestTargetsTable = "Requests"
	// RequestTargetsInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestTargetsInverseTable = "Requests"
	// RequestTargetsColumn is the table column denoting the requestTargets relation/edge.
	RequestTargetsColumn = "account_id"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEmail,
	FieldPassword,
	FieldType,
	FieldName,
	FieldAge,
	FieldGender,
	FieldAvatar,
	FieldIntroduction,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"account_id", "friend_id"}
	// RequestsPrimaryKey and RequestsColumn2 are the table columns denoting the
	// primary key for the requests relation (M2M).
	RequestsPrimaryKey = []string{"account_id", "request_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeActive   Type = "active"
	TypeInactive Type = "inactive"
	TypeAdmin    Type = "admin"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeActive, TypeInactive, TypeAdmin:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for type field: %q", _type)
	}
}

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale:
		return nil
	default:
		return fmt.Errorf("account: invalid enum value for gender field: %q", ge)
	}
}
