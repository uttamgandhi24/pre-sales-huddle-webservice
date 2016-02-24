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

type Mailer interface {
	GetEmailText() string
	GetEmailContext() string
}

func Notify(notificationPref NPType, mailer Mailer) {
	fmt.Println("NOTIFY")
	users := GetAllUsers()

	for _, user := range users {
		fmt.Println(user)
		for _, notification := range user.Notifications {
			if notification == notificationPref {
				fmt.Println("Send email for ", user.Email)
				emailMsg := EmailMessage{To: user.Email,
					Subject: NPTypeText[notificationPref] + mailer.GetEmailContext(),
					Body:    mailer.GetEmailText()}
				fmt.Println(emailMsg)
				SendEmail(emailMsg)
			}
		}
	}
}
