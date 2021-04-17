package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	// Routes
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
