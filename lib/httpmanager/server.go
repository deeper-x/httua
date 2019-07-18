package httpmanager

import (
	"fmt"
	"log"
	"net/http"
)

// Run server http
func Run() {
	fmt.Println("Server running...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1. Registration: PUT /register/<one-shot-unpredictable uri> - if true client is registered and userID is loaded in mem
		// 2A. IsRegistered: GET /id/<client-id> - if in mem, allow session
		// 2B. Deregistration: DELETE /id/<client-id> - if still in mem, delete it

		switch r.Method {
		case "GET":
			fmt.Println("#TODO - Is registered?")
		case "PUT":
			fmt.Println("#TODO - Registration")
		case "DELETE":
			fmt.Println("#TODO - Deregister. Good bye and see you soon")
		default:
			fmt.Println("#TODO - Method not allowed")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
