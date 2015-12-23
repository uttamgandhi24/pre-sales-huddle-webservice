package main

import (
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
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

/*func ProspectViewCriteriaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria := vars["criteria"]

	if err := json.NewEncoder(w).Encode(GetProspectsByCriteria(criteria)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}
*/
func ProspectAddHandler(w http.ResponseWriter, r *http.Request) {
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
}
func ProspectUpdateHandler(w http.ResponseWriter, r *http.Request) {

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

	if len(t.ProspectID) == 0 {
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
