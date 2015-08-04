package main

import (
	"fmt"
	"log"
	"strings"
)

type Participant struct {
	ProspectID    int    `bson:"ProspectID"`
	UserID        string `bson:"UserID"`
	Included      string `bson:"Included"`
	Participation string `bson:"Participation"`
}

func GetAllParticipants() (participants []Participant) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("select * from participants order by ProspectID asc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var participant Participant
		rows.Scan(&participant.ProspectID,
			&participant.UserID,
			&participant.Included,
			&participant.Participation)

		participants = append(participants, participant)
	}
	return participants
}

func GetParticipantByUserId(UserID string) (participants []Participant) {
	db := Connect()
	defer db.Close()
	str := "select * from participants where UserID='" + UserID + "'" + " Order by ProspectID asc"
	fmt.Println("getParticipantByUserID", str)
	rows, err := db.Query(str)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var participant Participant
		rows.Scan(&participant.ProspectID,
			&participant.UserID,
			&participant.Included,
			&participant.Participation)

		participants = append(participants, participant)
	}
	return participants
}

func GetParticipantByProspectId(ProspectID string) (participants []Participant) {
	db := Connect()
	defer db.Close()
	str := "select * from participants where ProspectID=" + ProspectID + " Order by UserID asc"
	rows, err := db.Query(str)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var participant Participant
		rows.Scan(&participant.ProspectID,
			&participant.UserID,
			&participant.Included,
			&participant.Participation)

		participants = append(participants, participant)
	}
	return participants
}

func (participant *Participant) Write() (err error) {
	db := Connect()
	defer db.Close()

	str := `INSERT INTO participants(
      ProspectID, UserID, Included, Participation)
      values(?,?,?,?)`
	stmt, err := db.Prepare(str)

	_, err = stmt.Exec(participant.ProspectID,
		participant.UserID, participant.Included, participant.Participation)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return err
}

func (participant *Participant) Update() (err error) {
	db := Connect()
	defer db.Close()

	if len(participant.UserID) == 0 {
		return err
	}
	if participant.ProspectID < 0 {
		return err
	}

	updateStr := ""
	if len(participant.Included) > 0 {
		updateStr = updateStr + " Included = '" + participant.Included + "',"
	}

	if len(participant.Participation) > 0 {
		updateStr = updateStr + " Participation = '" + participant.Participation + "'"
	}

	updateStr = strings.TrimSuffix(updateStr, ",")
	fmt.Println(updateStr)

	queryStr := "UPDATE participants Set" + updateStr + " where UserID=? AND ProspectID=?"

	stmt, err := db.Prepare(queryStr)

	_, err = stmt.Exec(participant.UserID, participant.ProspectID)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
