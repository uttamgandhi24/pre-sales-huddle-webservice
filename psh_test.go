package main

import(
    "net/http"
    "net/http/httptest"
    "testing"
        "github.com/gorilla/mux"
        //"bytes"
        "os"
)

func TestMain(m *testing.M) {
    //flag.Parse()
    router := mux.NewRouter()

    // Add handler functions for routes
    AddHandlers(router)

    // Dial into MongoDB database and get session handle
    session := ConnectDB()
    // TODO: Need to handle session.Close()

    gPshServer = PSHServer{router,session}

    http.Handle("/", &gPshServer)
    os.Exit(m.Run())
}

func TestProspectAll(t *testing.T) {
    req, _ := http.NewRequest("GET", "/prospect/all/", nil)
    w := httptest.NewRecorder()
    gPshServer.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("/prospect/all/ didn't return %v", http.StatusOK)
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
func TestParticipantAll(t *testing.T) {
    req, _ := http.NewRequest("GET", "/participant/all/", nil)
    w := httptest.NewRecorder()
    gPshServer.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("/participant/all/ didn't return %v", http.StatusOK)
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