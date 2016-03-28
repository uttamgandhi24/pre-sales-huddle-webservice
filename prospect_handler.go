package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func ProspectViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetAllProspects()); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}

func ProspectViewByProspectIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prospectid := vars["prospectid"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetProspectByProspectId(prospectid)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}

func ProspectAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var prospect Prospect
	err = json.Unmarshal(body, &prospect)

	if prospect.Name == "" {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}
	err = prospect.Write()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Notify(NPProspectCreated, prospect)
}

func ProspectConfCallAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var prospect Prospect
	err = json.Unmarshal(body, &prospect)

	if len(prospect.ProspectID) == 0 {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}
	err = prospect.AddConfCall()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Notify(NPCallScheduled, prospect)
}

func updateHandler(w http.ResponseWriter, r *http.Request, notificationPref NPType) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var prospect Prospect
	err = json.Unmarshal(body, &prospect)
	if err != nil {
		panic(err)
	}

	if len(prospect.ProspectID) == 0 {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		fmt.Println(prospect.ProspectID)
		return
	}
	fmt.Println("prospectupdate ", prospect)
	err = prospect.Update()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Notify(notificationPref, prospect)
}

func ProspectUpdateHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r, NPProspectUpdated)
}

func ProspectToClientHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r, NPProspectClient)
}

func ProspectToDeadHandler(w http.ResponseWriter, r *http.Request) {
	updateHandler(w, r, NPProspectDead)
}

