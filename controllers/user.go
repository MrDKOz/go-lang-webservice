package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// Bind this method to the userController
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Write back to the HTTP response with a string converted to a slice of bytes
	w.Write([]byte("Hello from the User Controller!"))

}

// Constructor function
// Returns a pointer to a userController
func newUserController() *userController {

	// Create the userController
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
