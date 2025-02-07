// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"

	// EdgeUser2room holds the string denoting the user2room edge name in mutations.
	EdgeUser2room = "user2room"
	// EdgeUser2problem holds the string denoting the user2problem edge name in mutations.
	EdgeUser2problem = "user2problem"

	// Table holds the table name of the user in the database.
	Table = "users"
	// User2roomTable is the table the holds the user2room relation/edge.
	User2roomTable = "rooms"
	// User2roomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	User2roomInverseTable = "rooms"
	// User2roomColumn is the table column denoting the user2room relation/edge.
	User2roomColumn = "user_id"
	// User2problemTable is the table the holds the user2problem relation/edge.
	User2problemTable = "problems"
	// User2problemInverseTable is the table name for the Problem entity.
	// It exists in this package in order to avoid circular dependency with the "problem" package.
	User2problemInverseTable = "problems"
	// User2problemColumn is the table column denoting the user2problem relation/edge.
	User2problemColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
