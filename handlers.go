package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AddHandlers(router *mux.Router) {
	router.HandleFunc("/connect/", ConnectHandler).Methods("POST")

	//Prospects
	router.HandleFunc("/prospect/all/", ProspectViewHandler).Methods("GET")
	router.HandleFunc("/prospect/prospectid/{prospectid}", ProspectViewByProspectIdHandler).Methods("GET")
	router.HandleFunc("/prospect/", ProspectAddHandler).Methods("POST")
	router.HandleFunc("/prospect/", ProspectUpdateHandler).Methods("PUT")
	router.HandleFunc("/prospect/confcall",
		ProspectConfCallAddHandler).Methods("POST")

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
	router.HandleFunc("/discussion/answer",
		DiscussionAnswerAddHandler).Methods("POST")

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
            X-CSRF-Token, Authentication`)
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	if r.URL.String() != "/connect/" && AuthenticateJWT(r.Header) == false {
		http.Error(w, "Unauthorized Access", http.StatusUnauthorized)
		return
	}
	// Lets Gorilla work
	server.router.ServeHTTP(w, r)
}

type Authentication struct {
	User  string `bson:"User",omitempty`
	Token string `bson:"Token",omitempty`
}

type GoogleResponse struct {
	iss   string `bson:"iss"`
	exp   int64  `bson:"exp"`
	email string `bson:"email"`
}

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(reqbody))
	var auth Authentication
	err = json.Unmarshal(reqbody, &auth)
	if err != nil {
		panic(err)
	}
	fmt.Println(auth)

	if len(auth.User) == 0 || len(auth.Token) == 0 {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	emailParts := strings.Split(auth.User, "@")
	if len(emailParts) != 2 || emailParts[1] != "synerzip.com" {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}
	request := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", auth.Token)
	resp, err := http.Get(request)
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return
	}
	respbody, err := ioutil.ReadAll(resp.Body)

	var googleResponse map[string]interface{}
	err = json.Unmarshal(respbody, &googleResponse)
	fmt.Println(googleResponse)
	if auth.User != googleResponse["email"] {
		fmt.Println("email mismatch")
		return
	}

	expireTime, _ := strconv.Atoi(googleResponse["exp"].(string))

	if time.Now().After(time.Unix(int64(expireTime), 0)) {
		fmt.Println("Token expired")
		return
	}

	token, err := generateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		panic(err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateJWT() (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err = token.SignedString([]byte("secret"))
	return tokenString, err
}
func AuthenticateJWT(header map[string][]string) bool {
	if header["Authentication"] == nil {
		return false
	}

	jwtString := header["Authentication"][0]
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		fmt.Println("Authenticate JWT failed", err)
		return false
	}
	return false
}
