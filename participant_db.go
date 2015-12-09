package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Participant struct {
	ProspectID    bson.ObjectId `bson:"ProspectID"`
	UserID        string        `bson:"UserID"`
	Included      string        `bson:"Included"`
	Participation string        `bson:"Participation"`
}

func GetAllParticipants() (participants []Participant) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	iter := collection.Find(nil).Iter()

	var participant Participant
	for iter.Next(&participant) {
		participants = append(participants, participant)
	}
	return participants
}

func GetParticipantByUserId(userID string) (participants []Participant) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	iter := collection.Find(Participant{UserID : userID}).Iter()
	var participant Participant
	for iter.Next(&participant) {
		participants = append(participants, participant)
	}
	return participants
}

func GetParticipantByProspectId(prospectID string) (participants []Participant) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	prospectIDHex := bson.ObjectIdHex(prospectID)
	iter := collection.Find(Participant{ProspectID : prospectIDHex}).Iter()
	var participant Participant
	for iter.Next(&participant) {
		participants = append(participants, participant)
	}
	return participants
}

func (participant *Participant) Write() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kParticipantsTable)

	err = collection.Insert(participant)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (participant *Participant) Update() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	collection.Update(Participant{UserID : participant.UserID,
		ProspectID : participant.ProspectID}, participant)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
