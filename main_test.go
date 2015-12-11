package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//flag.Parse()
	router := mux.NewRouter()

	// Add handler functions for routes
	AddHandlers(router)

	// Dial into MongoDB database and get session handle
	session := ConnectDB()
	defer session.Close()

	gPshServer = PSHServer{router, session}

	http.Handle("/", &gPshServer)
	runvalue := m.Run()
	CleanUp()
	os.Exit(runvalue)
}

func CleanUp() {
	ProspectTestCleanUp()
	ParticipantTestCleanUp()
	DiscussionTestCleanUp()
	UserTestCleanUp()
}
