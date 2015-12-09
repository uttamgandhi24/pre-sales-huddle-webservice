package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
  "gopkg.in/mgo.v2"
)

type PSHServer struct {
	router *mux.Router
	session *mgo.Session
}

// It's a global PSHServer object holding handle to gorilla router and mongodb
// session
var gPshServer PSHServer

func main() {
	// Create a new gorilla router
	router := mux.NewRouter()

	// Add handler functions for routes
	AddHandlers(router)

 // Dial into MongoDB database and get session handle
	session := ConnectDB()
	defer session.Close()

	gPshServer = PSHServer{router,session}

	http.Handle("/", &gPshServer)

	fmt.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
