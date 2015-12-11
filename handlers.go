package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AddHandlers(router *mux.Router) {
	router.HandleFunc("/prospect/all/", ProspectViewHandler).Methods("GET")
	router.HandleFunc("/prospect/", ProspectAddHandler).Methods("POST")
	router.HandleFunc("/prospect/", ProspectUpdateHandler).Methods("PUT")

	router.HandleFunc("/participant/all/", ParticipantViewHandler).Methods("GET")
	router.HandleFunc("/participant/userid/{userid}", ParticipantViewByUserId).
		Methods("GET")
	router.HandleFunc("/participant/prospectid/{id}",
		ParticipantViewByProspectId).Methods("GET")
	router.HandleFunc("/participant/", ParticipantAddHandler).Methods("POST")
	router.HandleFunc("/participant/", ParticipantUpdateHandler).
		Methods("PUT")

	router.HandleFunc("/discussion/all/", DiscussionViewHandler).Methods("GET")
	router.HandleFunc("/discussion/prospectid/{id}",
		DiscussionViewByProspectId).Methods("GET")
	router.HandleFunc("/discussion/", DiscussionAddHandler).Methods("POST")
	router.HandleFunc("/discussion/", DiscussionUpdateHandler).Methods("PUT")

	router.HandleFunc("/user/all/", UserViewHandler).Methods("GET")
	router.HandleFunc("/user/email/{email}", UserViewByEmail).Methods("GET")
	router.HandleFunc("/user/", UserAddHandler).Methods("POST")

}

func (server *PSHServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", `POST, GET, OPTIONS,
        	PUT, DELETE`)
		w.Header().Set("Access-Control-Allow-Headers",
			`Accept, Content-Type, Content-Length, Accept-Encoding,
            X-CSRF-Token, Authorization`)
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	server.router.ServeHTTP(w, r)
}
