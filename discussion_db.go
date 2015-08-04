package main

import (
	"log"
)

type Discussion struct {
	DiscussionID int    `bson:"DiscussionID"`
	ProspectID   int    `bson:"ProspectID"`
	UserID       string `bson:"UserID"`
	Query        string `bson:"Query"`
	Answer       string `bson:"Answer"`
}

func GetAllDiscussions() (discussions []Discussion) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("select * from discussions")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var discussion Discussion
		rows.Scan(&discussion.DiscussionID, &discussion.ProspectID,
			&discussion.UserID,
			&discussion.Query,
			&discussion.Answer)

		discussions = append(discussions, discussion)
	}
	return discussions
}
func GetDiscussionByProspectId(ProspectID string) (discussions []Discussion) {
	db := Connect()
	defer db.Close()
	str := "select * from discussions where ProspectID=" + ProspectID + " order by DiscussionID desc"
	rows, err := db.Query(str)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var discussion Discussion
		rows.Scan(&discussion.DiscussionID, &discussion.ProspectID,
			&discussion.UserID,
			&discussion.Query,
			&discussion.Answer)

		discussions = append(discussions, discussion)
	}
	return discussions
}

func (discussion *Discussion) Write() (err error) {
	db := Connect()
	defer db.Close()

	// insert
	str := `INSERT INTO discussions(
      ProspectID, UserID, Query, Answer)
      values(?,?,?,?)`
	stmt, err := db.Prepare(str)

	_, err = stmt.Exec(discussion.ProspectID,
		discussion.UserID, discussion.Query, discussion.Answer)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return err
}

func (discussion *Discussion) Update() (err error) {
	db := Connect()
	defer db.Close()

	if discussion.DiscussionID < 0 || len(discussion.Answer) == 0 {
		return err
	}

	stmt, err := db.Prepare("UPDATE discussions Set Answer = ? where DiscussionID=?")

	_, err = stmt.Exec(discussion.Answer, discussion.DiscussionID)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
