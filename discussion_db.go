package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Answer struct {
	AnswerStr string `bson:"AnswerStr"`
	UserID    string `bson:"UserID"`
}

type Discussion struct {
	DiscussionID bson.ObjectId `bson:"DiscussionID"`
	ProspectID   bson.ObjectId `bson:"ProspectID"`
	UserID       string        `bson:"UserID"`
	Query        string        `bson:"Query"`
	Answers      []Answer      `bson:"Answers"`
}

func GetAllDiscussions() (discussions []Discussion) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)

	iter := collection.Find(nil).Iter()
	if iter == nil {
		log.Fatal(kDiscussionsTable, " Iter nil")
		return
	}

	var discussion Discussion
	for iter.Next(&discussion) {
		discussions = append(discussions, discussion)
	}
	return
}
func GetDiscussionByProspectId(prospectID string) (discussions []Discussion) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)

	prospectid := bson.ObjectIdHex(prospectID)

	iter := collection.Find(bson.M{"ProspectID": prospectid}).Iter()
	var discussion Discussion
	for iter.Next(&discussion) {
		discussions = append(discussions, discussion)
	}
	return
}

func (discussion *Discussion) Write() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)

	discussion.DiscussionID = bson.NewObjectId()
	err = collection.Insert(discussion)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (discussion *Discussion) Update() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kDiscussionsTable)
	collection.Update(bson.M{"DiscussionID": discussion.DiscussionID}, discussion)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
