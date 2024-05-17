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
	Time       time.Time
	Identifier string
	Body       string
	TableNum   int
}

type Club struct {
	tables      int
	openingTime time.Time
	closingTime time.Time
	hourlyRate  int
	events      []Event
}

func NewClub(numTables int, openingTime, closingTime time.Time, hourlyRate int, events []Event) (*Club, error) {
	// if firstName == "" || lastName == "" || birthDate == "" {
	// 	return nil, errors.New("entered data can not be empty")
	// }

	return &Club{
		tables:      numTables,
		openingTime: openingTime,
		closingTime: closingTime,
		hourlyRate:  hourlyRate,
		events:      events,
	}, nil
}

func (club *Club) OutputClub() {
	fmt.Printf("Количество столов: %d\n", club.tables)
	fmt.Printf("Время начала работы: %s\n", club.openingTime.Format("09:41"))
	fmt.Printf("Время окончания работы: %s\n", club.closingTime.Format("09:41"))
	fmt.Printf("Стоимость часа: %d\n", club.hourlyRate)
}

type Gamer struct {
	Name string
}
