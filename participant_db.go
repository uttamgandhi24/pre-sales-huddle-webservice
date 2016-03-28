package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Participant struct {
	ProspectID        bson.ObjectId `bson:"ProspectID"`
	UserID            string        `bson:"UserID"`
	Included          string        `bson:"Included"`
	ParticipationRole string        `bson:"ParticipationRole"`
	AvailableDate     string        `bson:"AvailableDate"`
	Notes             string        `bson:"Notes"`
	ImageURL          string        `bson:"ImageURL"`
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

func GetParticipantByUserId(userID string) (participant Participant) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	collection.Find(bson.M{"UserID": userID}).One(&participant)
	return participant
}

func GetParticipantsByProspectId(prospectID string) (participants []Participant) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kParticipantsTable)
	prospectIDHex := bson.ObjectIdHex(prospectID)
	iter := collection.Find(bson.M{"ProspectID": prospectIDHex}).Iter()
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
	collection.Update(bson.M{"UserID": participant.UserID,
		"ProspectID": participant.ProspectID}, participant)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (participant Participant) GetEmailText(notificationPref NPType) (str string) {
	switch notificationPref {
	case NPSomeoneVolunteered:
		prospect := GetProspectByProspectId(participant.ProspectID.Hex())
		str = "Prospect Name: " + prospect.Name + "\r" +
			"ParticipationRole: " + participant.ParticipationRole + "\n"
	}
	return str
}

func (participant Participant) GetEmailContext(notificationPref NPType) (str string) {
	switch notificationPref {
	case NPSomeoneVolunteered:
		prospect := GetProspectByProspectId(participant.ProspectID.Hex())
		str = "A new volunteer for " + prospect.Name;
	}
	return str
}