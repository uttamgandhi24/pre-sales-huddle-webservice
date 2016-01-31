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
    "ProspectNotes":"Prospect Notes for the prospect",
    "ClientNotes":"Client Notes for prospect",
    "ConfCalls":[{"ConfDateStart":"2015-12-19T07:00",
			"ConfDateEnd":"2015-12-19T08:00",
			"ConfType":"PrepCall",
			"EnggFacilitator":"xyz@synerzip.com",
			"GoogleCalenderLink":"www.calendar.com"},
			{"ConfDateStart":"2015-12-20T07:00",
			"ConfDateEnd":"2015-12-20T08:00",
			"ConfType":"EnggCall",
			"GoogleCalenderLink":"www.calendar2.com",
			"EnggFacilitator":"xyzw@synerzip.com"}],
		"ProspectStatus":"NewlyCreated",
		"WebsiteURL":"www.synerzip.com"
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
	if prospect.ConfCalls[0].ConfType != "PrepCall" {
		t.Errorf("dummyProspect not added")
	}
	if prospect.ConfCalls[1].ConfDateStart != "2015-12-20T07:00" {
		t.Errorf("dummyProspect not added")
	}
	if prospect.ConfCalls[1].GoogleCalenderLink != "www.calendar2.com" {
		t.Errorf("dummyProspect not added")
	}
	if prospect.ProspectStatus != "NewlyCreated" {
		t.Errorf("dummyProspect not added")
	}
	if prospect.WebsiteURL != "www.synerzip.com" {
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
    "ProspectNotes":"Dummy Notes for the prospect",
    "ClientNotes":"Updated client Notes"
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
	if prospect.ClientNotes != "Updated client Notes" {
		t.Errorf("dummyProspect not updated")
	}
}

func TestProspectAddConfCall(t *testing.T) {
	// Get ProspectID for dummyProspect
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kProspectsTable)
	var prospect Prospect
	collection.Find(bson.M{"Name": "dummyProspect"}).One(&prospect)
	if prospect.Name != "dummyProspect" {
		t.Errorf("dummyProspect not added")
	}

	// Add confcall
	requestString := fmt.Sprintf(`{
    "ProspectID":"%v",
    "ConfCalls":[{"ConfDateStart":"2015-12-27T07:00",
			"ConfDateEnd":"2015-12-27T08:00",
			"ConfType":"PrepCall"},
			{"ConfDateStart":"2015-12-28T07:00",
			"ConfDateEnd":"2015-12-28T08:00",
			"ConfType":"EnggCall"}]
    }`, prospect.ProspectID.Hex())

	req, _ := http.NewRequest("POST", "/prospect/confcall",
		bytes.NewBuffer([]byte(requestString)))

	w := httptest.NewRecorder()
	gPshServer.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/prospect/confcall POST request didn't return %v", http.StatusOK)
	}

	//check if confcalls updated
	collection.Find(bson.M{"Name": "dummyProspect"}).One(&prospect)
	if prospect.ConfCalls[0].ConfType != "PrepCall" {
		t.Errorf("add confcall fail")
	}
	if prospect.ConfCalls[1].ConfDateStart != "2015-12-28T07:00" {
		t.Errorf("add confcall fail")
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
