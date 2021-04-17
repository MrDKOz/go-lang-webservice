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

// Return an array of pointers to the Users
func GetUsers() []*User {
	return users
}

// Add a new User to the list of users
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++

	users = append(users, &u)

	return u, nil
}
