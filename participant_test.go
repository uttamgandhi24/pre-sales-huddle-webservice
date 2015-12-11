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

func TestParticipantAdd(t *testing.T) {
	// Add dummyParticipant
	var reqStr = []byte(`{
    "ProspectID":"5665594d4ba30d74a3c3aaaa",
    "UserID":"abc@synerzip.com",
    "Included":"YES",
    "Participation":"YES"
  }`)
	req, _ := http.NewRequest("POST", "/participant/", bytes.NewBuffer(reqStr))
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/participant/ POST request didn't return %v", http.StatusOK)
	}

	//Connect to mongodb and check if dummyParticipant added
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kParticipantsTable)

	var participant Participant
	collection.Find(bson.M{"UserID": "abc@synerzip.com"}).One(&participant)
	if participant.UserID != "abc@synerzip.com" {
		t.Errorf("dummyParticipant not added")
	}
}

func TestParticipantUpdate(t *testing.T) {

	var reqStr = []byte(`{
    "ProspectID":"5665594d4ba30d74a3c3aaaa",
    "UserID":"abc@synerzip.com",
    "Included":"NO",
    "Participation":"YES"
  }`)

	req, _ := http.NewRequest("PUT", "/participant/",
		bytes.NewBuffer([]byte(reqStr)))

	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/participant/ PUT request didn't return %v", http.StatusOK)
	}

	//check if dummyPartcipant updated
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kParticipantsTable)

	var participant Participant
	collection.Find(bson.M{"UserID": "abc@synerzip.com"}).One(&participant)
	if strings.Compare(participant.Included, "NO") != 0 {
		t.Errorf("dummyParticipant not updated")
	}
}

func TestParticipantAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/participant/all/", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/participant/all/ didn't return %v", http.StatusOK)
	}
}

func TestParticipantByUserID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/participant/userid/abc@synerzip.com", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/participant/userid/{userid} didn't return %v", http.StatusOK)
	}
	var participant Participant
	json.Unmarshal([]byte(w.Body.String()), &participant)

	if strings.Compare(participant.UserID, "abc@synerzip.com") != 0 {
		t.Errorf("/participant/userid/{userid} didn't get the participant")
	}
}

func TestParticipantByProspectID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/participant/prospectid/5665594d4ba30d74a3c3aaaa", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/participant/prospectid/{prospectid} didn't return %v", http.StatusOK)
	}
	var participant []Participant
	json.Unmarshal([]byte(w.Body.String()), &participant)
	if strings.Compare(participant[0].UserID, "abc@synerzip.com") != 0 {
		t.Errorf("/participant/prospectid/{prospectid} didn't get the participant")
	}
}

func ParticipantTestCleanUp() {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	collection.Remove(bson.M{"UserID": "abc@synerzip.com"})
}
