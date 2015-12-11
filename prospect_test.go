package main

import (
	"bytes"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProspectAdd(t *testing.T) {
	// Add dummyProspect
	var reqStr = []byte(`{
    "Name":"dummyProspect",
    "TechStack":"Java",
    "Domain":"dummy",
    "TeamSize":3,
    "Notes":"Dummy Notes for the prospect"
  }`)
	req, _ := http.NewRequest("POST", "/prospect/", bytes.NewBuffer(reqStr))
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/prospect/ POST request didn't return %v", http.StatusOK)
	}

	//Connect to mongodb and check if dummyProspect added
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kProspectsTable)

	var prospect Prospect
	collection.Find(bson.M{"Name": "dummyProspect"}).One(&prospect)
	if prospect.Name != "dummyProspect" {
		t.Errorf("dummyProspect not added")
	}
}

func TestProspectUpdate(t *testing.T) {
	// Get ProspectID for dummyProspect
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kProspectsTable)
	var prospect Prospect
	collection.Find(bson.M{"Name": "dummyProspect"}).One(&prospect)
	if prospect.Name != "dummyProspect" {
		t.Errorf("dummyProspect not added")
	}

	// Update dummyProspect
	requestString := fmt.Sprintf(`{
    "ProspectID":"%v",
    "Name":"dummyProspect",
    "TechStack":"Java",
    "Domain":"dummy",
    "TeamSize":31,
    "Notes":"Dummy Notes for the prospect"
    }`, prospect.ProspectID.Hex())

	req, _ := http.NewRequest("PUT", "/prospect/",
		bytes.NewBuffer([]byte(requestString)))

	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/prospect/ PUT request didn't return %v", http.StatusOK)
	}

	//check if dummyProspect updated
	collection.Find(bson.M{"Name": "dummyProspect"}).One(&prospect)
	if prospect.TeamSize != 31 {
		t.Errorf("dummyProspect not updated")
	}
}

func TestProspectAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prospect/all/", nil)
	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/prospect/all/ didn't return %v", http.StatusOK)
	}
}

func ProspectTestCleanUp() {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kProspectsTable)
	collection.Remove(bson.M{"Name": "dummyProspect"})
}
