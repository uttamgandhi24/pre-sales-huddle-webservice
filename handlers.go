package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func AddHandlers() {
	h := mux.NewRouter()
	h.HandleFunc("/prospect/view/", ProspectViewHandler)
	h.HandleFunc("/prospect/view/{criteria}", ProspectViewCriteriaHandler)
	h.HandleFunc("/prospect/add/", ProspectAddHandler)
	h.HandleFunc("/prospect/update/", ProspectUpdateHandler)

	h.HandleFunc("/participant/add/", ParticipantAddHandler)
	h.HandleFunc("/participant/view/", ParticipantViewHandler)
	h.HandleFunc("/participant/view/userid/{userid}", ParticipantViewByUserId)
	h.HandleFunc("/participant/view/prospectid/{id:[0-9]+}", ParticipantViewByProspectId)
	h.HandleFunc("/participant/update/", ParticipantUpdateHandler)

	h.HandleFunc("/discussion/add/", DiscussionAddHandler)
	h.HandleFunc("/discussion/view/", DiscussionViewHandler)
	h.HandleFunc("/discussion/view/prospectid/{id:[0-9]+}", DiscussionViewByProspectId)
	h.HandleFunc("/discussion/update/", DiscussionUpdateHandler)
	//TODO remove comment for enabling DiscussionViewHTML
	//h.HandleFunc("/discussion/view/html/", DiscussionHTMLviewHandler)

	err := http.ListenAndServeTLS(":8080","server.pem", "server.key", h)

	fmt.Println("Listening on 8080....")
	if err!=nil {
		log.Fatal(err)
	}
}

func AuthenticateRequest(header map[string][]string, requestType string, requestKey string) bool {
	return true //TODO disabling authentication as of now, remove this to enable
	if header["Authentication"] == nil {
		return false
	}
	HMACValue := header["Authentication"][0]
	if HMACValue != ComputeHmac256(requestType, requestKey) {
		fmt.Println(ComputeHmac256(requestType, requestKey))
		return false
	}
	return true
}
