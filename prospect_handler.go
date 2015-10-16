package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// strings used for HMAC auth
const kProspectAdd string = "POST+/prospect/add/"
const kProspectView string = "GET+/prospect/view/"
const kProspectUpdate string = "POST+/prospect/update/"

const kProspectKey string = "PRESALES_PROSPECT_KEY"

func ProspectViewHandler(w http.ResponseWriter, r *http.Request) {
	if !AuthenticateRequest(r.Header, kProspectView, kProspectKey) {
		http.Error(w, "Authentication Error", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetAllProspects()); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}

func ProspectViewCriteriaHandler(w http.ResponseWriter, r *http.Request) {
	if !AuthenticateRequest(r.Header, kProspectView, kProspectKey) {
		http.Error(w, "Authentication Error", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	criteria := vars["criteria"]

	if err := json.NewEncoder(w).Encode(GetProspectsByCriteria(criteria)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}

func ProspectAddHandler(w http.ResponseWriter, r *http.Request) {
	if !AuthenticateRequest(r.Header, kProspectAdd, kProspectKey) {
		http.Error(w, "Authentication Error", http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t Prospect
	err = json.Unmarshal(body, &t)

	if t.Name == "" {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}
	err = t.Write()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProspectUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !AuthenticateRequest(r.Header, kProspectUpdate, kProspectKey) {
		http.Error(w, "Authentication Error", http.StatusUnauthorized)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t Prospect
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	if t.ProspectID <= 0 {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		fmt.Println(t.ProspectID)
		return
	}
	fmt.Println("prospectupdate ", t)
	err = t.Update()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
