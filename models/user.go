package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	// Hold a slice of pointers to Users
	users []*User

	// Not defining a datatype as it's implied as an int
	nextID = 1
)
