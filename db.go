package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

const kMongoConnectURL string = "localhost"
const kPreSalesDB string = "presales_huddle"
const kProspectsTable string = "prospects"
const kDiscussionsTable string = "discussions"
const kParticipantsTable string = "participants"
const kUsersTable string = "users"

func ConnectDB() (session *mgo.Session) {
	session, err := mgo.Dial(kMongoConnectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}
