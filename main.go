package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Create a new gorilla router
	router := mux.NewRouter()

	// Add handler functions for routes
	AddHandlers(router)

	http.Handle("/", &PSHServer{router})

	fmt.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
