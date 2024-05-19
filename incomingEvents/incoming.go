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

func Revenue(club *objects.Club) {
	tables := club.GetTablesRevenue()
	for key, value := range tables {
		fmt.Printf("%d %d\n", key, value.Revenue())
	}

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
		club.AddGamer(event.Body(), event.TableNum())

		//запуск подсчета времени и стоимости посещения стола
		tableRevenue := club.GetTablesRevenue()
		currTable, exists := tableRevenue[event.TableNum()]
		if exists {
			currTable.SetStartTime(event.Time())
			tableRevenue[event.TableNum()] = currTable
			club.SetTablesRevenue(tableRevenue)
		} else {
			table, err := objects.NewTable(event.TableNum(), event.Time(), event.Time())
			if err != nil {
				fmt.Println("id2")
				return err
			}
			m := make(map[int]objects.Table)
			m[event.TableNum()] = *table
			club.SetTablesRevenue(m)
		}

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
		club.RemoveFromWaitList(event.Body())
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
		club.RemoveFromWaitList(event.Body())
	}
	return nil
}

func Id4(event *objects.Event, club *objects.Club) error {
	if !club.IsVisitorInClub(event.Body()) {
		//если посетителя нет в клубе
		clientUnknown, err := objects.NewEvent(event.Time(), "13", "ClientUnknown", 0)
		if err != nil {
			fmt.Println("id4")
			return err
		}
		clientUnknown.PrintEvent()
		return nil
	}
	gamer := event.Body()
	table := club.GetGamerTable(gamer)
	club.RemoveTable(table)
	club.RemoveGamer(gamer)
	club.RemoveVisitor(gamer)

	//окончание подсчета стоимости стола
	tableRevenue := club.GetTablesRevenue()
	t := tableRevenue[table]
	t.SetEndTime(event.Time())
	t.UpdateOccupiedTime()
	t.UpdateDuration()
	t.UpdateRevenue(club.HourlyRate())
	tableRevenue[table] = t
	club.SetTablesRevenue(tableRevenue)

	if club.GetWaitListLength() != 0 {
		err := Id12(event, club, table)
		if err != nil {
			return err
		}
	}
	return nil
}

func Id11(club *objects.Club) error {
	//все посетители уходят при закрытии
	clubCloses := club.ClubCloses()
	for i := 0; i < len(*clubCloses); i++ {
		eventClose, err := objects.NewEvent(club.ClosingTime(), "11", (*clubCloses)[i], 0)
		if err != nil {
			fmt.Println("id11")
			return err
		}
		//окончание подсчета стоимости стола
		gamer := (*clubCloses)[i]
		table := club.GetGamerTable(gamer)
		tableRevenue := club.GetTablesRevenue()
		t := tableRevenue[table]
		t.SetEndTime(club.ClosingTime())
		t.UpdateOccupiedTime()
		t.UpdateDuration()
		t.UpdateRevenue(club.HourlyRate())
		tableRevenue[table] = t
		club.SetTablesRevenue(tableRevenue)
		eventClose.PrintEvent()

	}
	return nil
}

func Id12(event *objects.Event, club *objects.Club, freeTable int) error {
	client := club.GetClientFromWaitList(0)
	clientFromWaitList, err := objects.NewEvent(event.Time(), "12", client, freeTable)
	if err != nil {
		fmt.Println("id12")
		return err
	}
	clientFromWaitList.PrintEvent()
	club.AddTable(freeTable, client)
	club.AddGamer(client, freeTable)

	//запуск подсчета времени и стоимости посещения стола
	tableRevenue := club.GetTablesRevenue()
	currTable, exists := tableRevenue[freeTable]
	if exists {
		currTable.SetStartTime(event.Time())
		tableRevenue[freeTable] = currTable
		club.SetTablesRevenue(tableRevenue)
	} else {
		table, err := objects.NewTable(freeTable, event.Time(), event.Time())
		if err != nil {
			fmt.Println("id12")
			return err
		}
		m := make(map[int]objects.Table)
		m[freeTable] = *table
		club.SetTablesRevenue(m)
	}

	return nil
}
