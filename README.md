
# Simple GoLang Webservice
This is a repo containing a simple web service written in Go. This was made using the PluralSight course "Go: Getting Started" by Mike Van Sickle.

This course (and project) covers the following basic parts of Go:
 - Working with data types
 - Creating Functions and Methods
 - Controlling Program flow

The web service allows you to:
 - Fetch all Users
 - Fetch a specific User
 - Create a new User
 - Update an existing User
 - Delete an existing User

If you wish to run the project, follow these steps:

 - Ensure you have Go installed on your machine
	 - You can run `go version` in your terminal/command prompt to see if this is the case
	 - *if you don't* you can download it from [golang.org](https://golang.org/)
 - Clone the repo
 - Navigate to the root directory using terminal/command prompt
 - Run `go build .`
 - You're now running the web service on port 3000!

To use the web service you can send the following requests to "http://localhost:3000"
| Method | URL | Request Body | Description |
|--|--|--|--|
| GET | {rootUrl}/users | N/A | Returns all Users
| GET | {rootUrl}/users/1 | N/A | Returns the User with the ID of 1
| POST | {rootUrl}/users | `{"FirstName":"John","LastName":"Smith"}` | Creates a new User with the provided details
| PUT | {rootUrl}/users/1 | `{"ID":1,"FirstName":"Johnny","LastName":"Johnson"}` | Updates the User with matching ID with the provided details. *Note* User ID in the request, must match the ID in the URL
| DELETE | {rootUrl}/users/1 | N/A | Deletes the User with the matching ID
