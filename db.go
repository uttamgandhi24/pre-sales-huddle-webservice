package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

const kMongoConnectURL string = "localhost"
const kPreSalesDB string = "presales_huddle"
const kProspectsTable string = "prospects"
const kDiscussionsTable string = "discussions"
const kParticipantsTable string = "participants"
const kUsersTable string = "users"

func ConnectDB() (session *mgo.Session) {
	var mongoConnectURL string
	if mongoConnectURL = os.Getenv("PRESALES_MONGO_CONNECT_URL"); len(mongoConnectURL) == 0 {
		mongoConnectURL = kMongoConnectURL
	}
	session, err := mgo.Dial(mongoConnectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}
