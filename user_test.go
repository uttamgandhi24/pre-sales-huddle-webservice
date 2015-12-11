package main

import (
	"bytes"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserAdd(t *testing.T) {
	// Add dummyUser
	var reqStr = []byte(`{
    "Email":"abc@synerzip.com",
    "Role":"Engineer"
  }`)
	req, _ := http.NewRequest("POST", "/user/", bytes.NewBuffer(reqStr))
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/user/ POST request didn't return %v", http.StatusOK)
	}

	//Connect to mongodb and check if dummyParticipant added
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kUsersTable)

	var user User
	collection.Find(bson.M{"Email": "abc@synerzip.com"}).One(&user)
	if strings.Compare(user.Role, "Engineer") != 0 {
		t.Errorf("dummyUser not added")
	}
}

func TestUserAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user/all/", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/user/all/ didn't return %v", http.StatusOK)
	}
}

func TestUserByEmail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user/email/abc@synerzip.com", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/user/email/{id} didn't return %v", http.StatusOK)
	}
	var user User
	json.Unmarshal([]byte(w.Body.String()), &user)
	if strings.Compare(user.Role, "Engineer") != 0 {
		t.Errorf("/user/email/{id} didn't get the user")
	}
}

func UserTestCleanUp() {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kUsersTable)
	collection.Remove(bson.M{"Email": "abc@synerzip.com"})
}
