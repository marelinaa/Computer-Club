package incomingEvents

import (
	"errors"
	"example.com/computer-club/objects"
	"fmt"
	"regexp"
)

func IsValidClientName(name string) bool {
	// Используем регулярное выражение для проверки соответствия имени клиента заданным условиям
	regex := regexp.MustCompile(`^[a-z0-9_-]+$`)
	return regex.MatchString(name)
}

func Id1(event *objects.Event, club *objects.Club) error {
	//если неподходящее имя клиента
	if !IsValidClientName(event.Body()) {
		return errors.New(event.EventToString())
	}
	//если пришел в нерабочее время
	if event.Time().Before(club.OpeningTime()) || event.Time().After(club.ClosingTime()) {
		notOpenEvent, err := objects.NewEvent(event.Time(), "13", "NotOpenYet", 0)
		if err != nil {
			fmt.Println("id1")
			return err
		}
		notOpenEvent.PrintEvent()
	} else if !club.IsVisitorInClub(event.Body()) {
		//если клиента не было в клубе
		club.AddVisitor(event.Body())
	} else {
		//если клиент был в клубе
		club.RemoveVisitor(event.Body())
		alreadyVisited, err := objects.NewEvent(event.Time(), "13", "YouShallNotPass", 0)
		if err != nil {
			fmt.Println("id1")
			return err
		}
		alreadyVisited.PrintEvent()
	}

	return nil
}

func Id2(event *objects.Event, club *objects.Club) error {
	if event.TableNum() > club.Tables() {
		return errors.New(event.EventToString())
	}

	if !club.IsVisitorInClub(event.Body()) {
		//если посетителя нет в клубе
		clientUnknown, err := objects.NewEvent(event.Time(), "13", "ClientUnknown", 0)
		if err != nil {
			fmt.Println("id2")
			return err
		}
		clientUnknown.PrintEvent()
	} else if club.WhoUsesTable(event.TableNum()) == "" {
		//если стол не занят
		club.AddTable(event.TableNum(), event.Body())
	} else if club.WhoUsesTable(event.TableNum()) != "" {
		//если стол занят
		placeIsBusy, err := objects.NewEvent(event.Time(), "13", "PlaceIsBusy", 0)
		if err != nil {
			fmt.Println("id2")
			return err
		}
		placeIsBusy.PrintEvent()
	}
	return nil
}

func Id3(event *objects.Event, club *objects.Club) error {
	club.AddToWaitList(event.Body())
	//Если в очереди ожидания клиентов больше, чем общее число столов
	if club.GetWaitListLength() > club.Tables() {
		clientLeft, err := objects.NewEvent(event.Time(), "11", event.Body(), 0)
		if err != nil {
			fmt.Println("id2")
			return err
		}
		clientLeft.PrintEvent()
		return nil
	}

	//Если в клубе есть свободные столы
	if club.GetCurrTableCount() < club.Tables() {
		canWaitNoLonger, err := objects.NewEvent(event.Time(), "13", "ICanWaitNoLonger!", 0)
		if err != nil {
			fmt.Println("id2")
			return err
		}
		canWaitNoLonger.PrintEvent()
	}
	return nil
}
