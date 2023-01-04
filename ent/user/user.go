// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// EdgeArticles holds the string denoting the articles edge name in mutations.
	EdgeArticles = "articles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ArticlesTable is the table that holds the articles relation/edge.
	ArticlesTable = "articles"
	// ArticlesInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticlesInverseTable = "articles"
	// ArticlesColumn is the table column denoting the articles relation/edge.
	ArticlesColumn = "author"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldLevel,
	FieldBio,
	FieldImage,
	FieldCreatedAt,
	FieldPoints,
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
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultLevel holds the default value on creation for the "level" field.
	DefaultLevel uint32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultPoints holds the default value on creation for the "points" field.
	DefaultPoints uint32
)