package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	Email         string   `bson:"Email"`
	Role          string   `bson:"Role"`
	Notifications NPArray `bson:"Notifications"`
}

func GetAllUsers() (users []User) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kUsersTable)
	iter := collection.Find(bson.M{}).Iter()

	var user User
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users
}

func GetUserByEmail(email string) (user User) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kUsersTable)
	collection.Find(bson.M{"Email": email}).One(&user)
	return
}

func (user *User) Write() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()
	collection := session.DB(kPreSalesDB).C(kUsersTable)

	err = collection.Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (user *User) UpdateNotification() (err error) {
	session := gPshServer.session.Copy()
	defer session.Close()

	collection := session.DB(kPreSalesDB).C(kUsersTable)

	// Add new call to conf call array
	err = collection.Update(bson.M{"Email": user.Email},
		bson.M{"$set": bson.M{"Notifications": user.Notifications}})
	if err != nil {

		log.Fatal("update notification error", err)
	}
	return err
}

func (user User) IsUserInterestedInNotification(notification NPType) (bool) {
	if user.Notifications.HasNotification(notification) {
		return true
	}
	return false
}
