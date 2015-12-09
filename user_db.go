package main

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	Email string `bson:"Email"`
	Role  string `bson:"Role"`
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
	collection.Find(User{Email: email}).One(&user)
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
