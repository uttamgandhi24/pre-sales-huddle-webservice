package main

import (
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDiscussionAdd(t *testing.T) {
	// Add dummyDiscussion
	var reqStr = []byte(`{
    "DiscussionID":"5665594d4ba30d74a3c3ac83",
    "ProspectID":"5665594d4ba30d74a3c3ac83",
    "UserID":"abc@synerzip.com",
    "Query":"Simple Question"
  }`)
	req, _ := http.NewRequest("POST", "/discussion/", bytes.NewBuffer(reqStr))
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/discussion/ POST request didn't return %v", http.StatusOK)
	}

	//Connect to mongodb and check if dummyDiscussion added
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)

	var discussion Discussion
	collection.Find(bson.M{"UserID": "abc@synerzip.com"}).One(&discussion)
	if discussion.UserID != "abc@synerzip.com" {
		t.Errorf("dummyDiscussion not added")
	}
}

func TestDiscussionUpdate(t *testing.T) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)
	var discussion Discussion
	collection.Find(bson.M{"UserID": "abc@synerzip.com"}).One(&discussion)
	if discussion.UserID != "abc@synerzip.com" {
		t.Errorf("dummyDiscussion not found")
	}

	// Update dummyProspect
	requestString := fmt.Sprintf(`{
    "DiscussionID":"%v",
    "ProspectID":"5665594d4ba30d74a3c3ac83",
    "UserID":"abc@synerzip.com",
    "Query":"Simple Question",
    "Answer":"Simple Answer"
    }`, discussion.DiscussionID.Hex())

	req, _ := http.NewRequest("PUT", "/discussion/",
		bytes.NewBuffer([]byte(requestString)))

	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/discussion/ PUT request didn't return %v", http.StatusOK)
	}

	//check if dummyDiscussion updated
	collection.Find(bson.M{"UserID": "abc@synerzip.com"}).One(&discussion)
	if strings.Compare(discussion.Answer, "Simple Answer") != 0 {
		t.Errorf("dummyDiscussion not updated")
	}
}

func TestDiscussionAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/discussion/all/", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/discussion/all/ didn't return %v", http.StatusOK)
	}
}

func DiscussionTestCleanUp() {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)
	collection.Remove(bson.M{"UserID": "abc@synerzip.com"})
}
