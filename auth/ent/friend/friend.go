// Code generated by ent, DO NOT EDIT.

package friend

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the friend type in the database.
	Label = "friend"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldFriendID holds the string denoting the friend_id field in the database.
	FieldFriendID = "friend_id"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeFriend holds the string denoting the friend edge name in mutations.
	EdgeFriend = "friend"
	// Table holds the table name of the friend in the database.
	Table = "Friends"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "Friends"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "Accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
	// FriendTable is the table that holds the friend relation/edge.
	FriendTable = "Friends"
	// FriendInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	FriendInverseTable = "Accounts"
	// FriendColumn is the table column denoting the friend relation/edge.
	FriendColumn = "friend_id"
)

// Columns holds all SQL columns for friend fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAccountID,
	FieldFriendID,
}

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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
