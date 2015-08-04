package main

import (
	"github.com/gorilla/mux"
	"net/http"
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
	h.HandleFunc("/discussion/view/html/", DiscussionHTMLviewHandler)

	http.ListenAndServe(":8080", h)
}
