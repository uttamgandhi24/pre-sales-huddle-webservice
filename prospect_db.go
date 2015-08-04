package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Prospect struct {
	ProspectID      int    `bson:"ProspectID",omitempty`
	Name            string `bson:"Name",omitempty`
	ConfDateStart   string `bson:"ConfDate",omitempty`
	ConfDateEnd     string `bson:"ConfDate",omitempty`
	TechStack       string `bson:"TechStack",omitempty`
	Domain          string `bson:"Domain",omitempty`
	DesiredTeamSize int    `bson:"DesiredTeamSize",omitempty`
	Notes           string `bson:"Notes",omitempty`
	SalesID         string `bson:"SalesID",omitempty`
	CreateDate      string `bson:"CreateDate",omitempty`
	StartDate       string `bson:"StartDate",omitempty`
	BUHead          string `bson:"BUHead",omitempty`
	TeamSize        int    `bson:"TeamSize",omitempty`
}

func GetAllProspects() (prospects []Prospect) {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("select * from prospects order by ProspectID asc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var prospect Prospect
		rows.Scan(&prospect.ProspectID, &prospect.Name, &prospect.ConfDateStart,
			&prospect.ConfDateEnd,
			&prospect.TechStack,
			&prospect.Domain,
			&prospect.DesiredTeamSize,
			&prospect.Notes,
			&prospect.SalesID,
			&prospect.CreateDate,
			&prospect.StartDate,
			&prospect.BUHead,
			&prospect.TeamSize)

		prospects = append(prospects, prospect)
	}
	return prospects
}

func GetProspectsByCriteria(criteria string) (prospects []Prospect) {
	db := Connect()
	defer db.Close()
	fmt.Println("getProspectsByCriteria", criteria)
	predicates := strings.Split(criteria, "&")
	predicate_str := ""
	for _, s := range predicates {
		values := strings.Split(s, ":")
		predicate_str = predicate_str + values[0] + " LIKE '%" + values[1] + "%' " + "AND "
	}
	predicate_str = strings.TrimSuffix(predicate_str, "AND ")

	fmt.Println(predicate_str)
	str := "select * from prospects where " + predicate_str
	rows, err := db.Query(str)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var prospect Prospect
		rows.Scan(&prospect.ProspectID,
			&prospect.Name,
			&prospect.ConfDateStart,
			&prospect.ConfDateEnd,
			&prospect.TechStack,
			&prospect.Domain,
			&prospect.DesiredTeamSize,
			&prospect.Notes,
			&prospect.SalesID,
			&prospect.CreateDate,
			&prospect.StartDate,
			&prospect.BUHead,
			&prospect.TeamSize)
		prospects = append(prospects, prospect)
	}
	return prospects
}

func (prospect *Prospect) Write() (err error) {
	db := Connect()
	defer db.Close()

	// insert
	str := `INSERT INTO prospects(
      Name, ConfDateStart, ConfDateEnd, TechStack, Domain, DesiredTeamSize,
      Notes, SalesID, CreateDate, StartDate, BUHead, TeamSize)
      values(?,?,?,?,?,?,?,?,?,?,?,?)`
	stmt, err := db.Prepare(str)

	_, err = stmt.Exec(prospect.Name,
		prospect.ConfDateStart,
		prospect.ConfDateEnd,
		prospect.TechStack,
		prospect.Domain,
		prospect.DesiredTeamSize,
		prospect.Notes,
		prospect.SalesID,
		prospect.CreateDate,
		prospect.StartDate,
		prospect.BUHead,
		prospect.TeamSize)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return err
}

func (prospect *Prospect) Update() (err error) {
	db := Connect()
	defer db.Close()

	if prospect.ProspectID < 0 {
		return err
	}

	updateStr := ""
	if len(prospect.Name) > 0 {
		updateStr = updateStr + " Name = '" + prospect.Name + "',"
	}

	if len(prospect.ConfDateStart) > 0 {
		updateStr = updateStr + " ConfDateStart = '" + prospect.ConfDateStart + "',"
	}

	if len(prospect.ConfDateEnd) > 0 {
		updateStr = updateStr + " ConfDateEnd = '" + prospect.ConfDateEnd + "',"
	}

	if len(prospect.TechStack) > 0 {
		updateStr = updateStr + " TechStack = '" + prospect.TechStack + "',"
	}

	if len(prospect.Domain) > 0 {
		updateStr = updateStr + " Domain = '" + prospect.Domain + "',"
	}

	if prospect.DesiredTeamSize > 0 {
		updateStr = updateStr + " DesiredTeamSize = '" + strconv.Itoa(prospect.DesiredTeamSize) + "',"
	}

	if len(prospect.Notes) > 0 {
		updateStr = updateStr + " Notes = '" + prospect.Notes + "',"
	}

	if len(prospect.SalesID) > 0 {
		updateStr = updateStr + " SalesID = '" + prospect.SalesID + "',"
	}

	if len(prospect.StartDate) > 0 {
		updateStr = updateStr + " StartDate = '" + prospect.StartDate + "',"
	}

	if len(prospect.BUHead) > 0 {
		updateStr = updateStr + " BUHead = '" + prospect.BUHead + "',"
	}

	if prospect.TeamSize > 0 {
		updateStr = updateStr + " TeamSize = '" + strconv.Itoa(prospect.TeamSize) + "',"
	}
	updateStr = strings.TrimSuffix(updateStr, ",")
	fmt.Println(updateStr)

	queryStr := "UPDATE prospects Set" + updateStr + " where ProspectID=?"
	fmt.Println(queryStr)

	stmt, err := db.Prepare(queryStr)
	_, err = stmt.Exec(prospect.ProspectID)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
