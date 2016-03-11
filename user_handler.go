package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func UserViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetAllUsers()); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}
func UserAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t User
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
func UserViewByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetUserByEmail(email)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}
func UserNotificationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In UserNotificationHandler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var user User
	err = json.Unmarshal(body, &user)
	fmt.Println("In user", user.Email)
	if len(user.Email) == 0 {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}
	err = user.UpdateNotification()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
