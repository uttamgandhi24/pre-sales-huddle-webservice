package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
)

type PSHServer struct {
	router  *mux.Router
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

	gPshServer = PSHServer{router, session}

	http.Handle("/", &gPshServer)
	//TODO: use -indexDir flag to take input from user
	http.Handle("/presaleshuddle/", http.StripPrefix("/presaleshuddle/",
		http.FileServer(http.Dir("../../PreSales-Huddle/app/"))))

	//TODO: Take -port input from user
	fmt.Println("Listening on 8080")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println(err)
	}
}
