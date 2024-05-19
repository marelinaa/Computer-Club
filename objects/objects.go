package objects

import (
	"fmt"
	"time"
)

type Table struct {
	Number   int
	Revenue  int
	Duration int
}

type Event struct {
	time       time.Time
	identifier string
	body       string
	tableNum   int
}

func NewEvent(time time.Time, identifier, body string, tableNum int) (*Event, error) {
	return &Event{
		time:       time,
		identifier: identifier,
		body:       body,
		tableNum:   tableNum,
	}, nil
}

func (event *Event) Time() time.Time {
	return event.time
}

func (event *Event) Identifier() string {
	return event.identifier
}

func (event *Event) Body() string {
	return event.body
}

func (event *Event) PrintEvent() {
	fmt.Printf("%s %s %s\n", event.time.Format("15:04"), event.identifier, event.body)
}

type Club struct {
	tables      int
	openingTime time.Time
	closingTime time.Time
	hourlyRate  int
	//events      []Event
}

func NewClub(numTables int, openingTime, closingTime time.Time, hourlyRate int) (*Club, error) {
	// if firstName == "" || lastName == "" || birthDate == "" {
	// 	return nil, errors.New("entered data can not be empty")
	// }

	return &Club{
		tables:      numTables,
		openingTime: openingTime,
		closingTime: closingTime,
		hourlyRate:  hourlyRate,
		//events:      events,
	}, nil
}

func (club *Club) OpeningTime() time.Time {
	return club.openingTime
}

func (club *Club) ClosingTime() time.Time {
	return club.closingTime
}

func (club *Club) PrintClub() {
	fmt.Printf("Количество столов: %d\n", club.tables)
	fmt.Printf("Время начала работы: %s\n", club.openingTime.Format("15:04"))
	fmt.Printf("Время окончания работы: %s\n", club.closingTime.Format("15:04"))
	fmt.Printf("Стоимость часа: %d\n", club.hourlyRate)
}

type Gamer struct {
	Name string
}
