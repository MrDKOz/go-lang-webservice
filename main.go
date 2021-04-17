package main

import (
	"example-webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	// Listen on port 3000, and use default handler
	http.ListenAndServe(":3000", nil)
}
