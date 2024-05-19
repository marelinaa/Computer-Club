package incomingEvents

import (
	"example.com/computer-club/objects"
	"fmt"
)

func Id1(event *objects.Event, club *objects.Club) {
	if event.Time().Before(club.OpeningTime()) || event.Time().After(club.ClosingTime()) {
		notOpenEvent, err := objects.NewEvent(event.Time(), event.Identifier(), "NotOpenYet", 0)
		if err != nil {
			fmt.Println("aaaaAAAAAAAAAINCOMING")
			return
		}
		notOpenEvent.PrintEvent()
	}

}
