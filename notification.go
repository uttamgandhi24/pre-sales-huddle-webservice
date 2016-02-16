package main

import (
	"fmt"
)

type NPType int

const (
	NPEveryProspect NPType = iota
	NPRelevantProspect
	NPProspectCreated
	NPProspectUpdated
)

var NPTypeText = map[NPType]string{
	NPProspectCreated: "A new Prospect Added by ",
}

func Notify(notificationPref NPType, prospect Prospect) {
	fmt.Println("NOTIFY")
	users := GetAllUsers()

	for _, user := range users {
		fmt.Println(user)
		for _, notification := range user.Notifications {
			if notification == notificationPref {
				fmt.Println("Send email for ", user.Email)
				emailMsg := EmailMessage{To: user.Email,
					Subject: NPTypeText[notificationPref] + prospect.SalesID,
					Body:    prospect.MarshalEmail()}
				fmt.Println(emailMsg)
				SendEmail(emailMsg)
			}
		}
	}
}
