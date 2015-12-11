package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func ParticipantViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetAllParticipants()); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}
func ParticipantAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t Participant
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	err = t.Write()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ParticipantViewByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetParticipantByUserId(userid)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}
func ParticipantViewByProspectId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prospectid := vars["id"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetParticipantsByProspectId(prospectid)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}

func ParticipantUpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t Participant
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	err = t.Update()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
