package main

import (
	"fmt"
)

type NPType int

const (
	NPEveryProspect      NPType = iota
	NPRelevantProspect          // 1
	NPProspectCreated           // 2
	NPProspectUpdated           // 3
	NPQuestionPosted            // 4
	NPQuestionAnswered          // 5
	NPSomeoneVolunteered        // 6
	NPCallScheduled             // 7
	NPProspectDead              // 8
	NPProspectClient            // 9
)

var NPTypeText = map[NPType]string{
	NPProspectCreated: "A new Prospect Added by ",
}

type Mailer interface {
	GetEmailText(notificationPref NPType) string
	GetEmailContext(notificationPref NPType) string
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
					Subject: NPTypeText[notificationPref] + mailer.GetEmailContext(notificationPref),
					Body:    mailer.GetEmailText(notificationPref)}
				fmt.Println(emailMsg)
				// SendEmail(emailMsg)
			}
		}
	}
}
