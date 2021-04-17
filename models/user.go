package models

import (
	"errors"
	"fmt"
)

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

// Return a slice of pointers to the Users
func GetUsers() []*User {
	return users
}

// Add a new User to the list of users
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include an id, or it must be 0")
	}

	// Assign an ID to the new user, then increment nextID
	u.ID = nextID
	nextID++

	// Append the User pointer to the slice
	users = append(users, &u)

	// Return the User, no error occurred
	return u, nil
}

// Checks all stored Users for a matching ID
// If found: Returns a pointer to the matched User
// If not found: Returns an empty user, and a formatted error
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// Accepts a User struct, and replaces the matching place in
// the slice, with the update User details
// If a User is found, return the user, and no error
// If a User is not found, return an empty User, and a formatted error
func UpdateUser(u User) (User, error) {
	for i, candiate := range users {
		if candiate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// Accepts a User ID, and removes it from the stored Users
// If User is found, create a new slice, without the matched User
// If User was not found, return a formatted error
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' was not found", id)
}
